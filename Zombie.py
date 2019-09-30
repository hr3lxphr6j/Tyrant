#!/usr/bin/env python
import argparse
import functools
import os
import re
import sqlite3
import subprocess
import time
from queue import Queue

_name_ = 'Zombie'
_version_ = '0.1'

ffmpeg_args = '-y -c:v libx265 -crf 22 -preset faster -c:a aac -b:a 192k -map_metadata -1 -max_muxing_queue_size 4096'


def parse_args():
    parser = argparse.ArgumentParser()
    parser.add_argument('--ffmpeg-args',
                        help='ffmpeg options without input and output, like "-c:v libx265"',
                        default=ffmpeg_args)
    parser.add_argument('-r', '--recursive',
                        action="store_true",
                        help='recurse into directories',
                        default=False)
    parser.add_argument('-d', '--delete',
                        action="store_true",
                        help='delete origin file after process success',
                        default=False)
    parser.add_argument('--target-exts',
                        help='target file exts, split with ",", like: mp4,mkv',
                        dest='exts',
                        default=['.mp4', '.webm', '.mkv', '.flv'],
                        type=lambda str: ["."+ext.strip() for ext in str.split(',')])
    parser.add_argument('--debug',
                        action="store_true",
                        help='only show command without execute',
                        default=False)
    parser.add_argument('file', nargs='+', help='target file or directory')
    return parser.parse_args()


def files(files, recursive=False):
    queue = Queue()
    # split file and dir
    for file in files:
        if os.path.isfile(file):
            yield os.path.split(file)[::-1]
        if os.path.isdir(file):
            queue.put(file)
    # parse dir
    while not queue.empty():
        root = queue.get()
        for f in os.listdir(root):
            full_name = os.path.join(root, f)
            if os.path.isdir(full_name) and recursive:
                queue.put(full_name)
            yield f, root


def get_description_metadata(file):
    res = subprocess.run([
        'ffmpeg',
        '-i', file,
        '-f', 'ffmetadata',
        '-loglevel', 'quiet',
        '-'
    ], capture_output=True)
    assert res.returncode == 0, 'file: {}, failed to get metadata'.format(file)
    for data in res.stdout.split(b'\n'):
        if data.startswith(b'description='):
            return data.split(b'=')[-1].decode()


def call(command):
    subprocess.call(command)


def record(fn):
    @functools.wraps(fn)
    def wrapper(*args, **kw):
        debug = kw.get('debug') or False
        if debug:
            return fn(*args, **kw)
        input_file = kw.get('input')
        output_file = kw.get('output')
        start_time = int(time.time())
        result = fn(*args, **kw)
        finish_time = int(time.time())
        with sqlite3.connect('report.db') as db:
            db.execute('''CREATE TABLE IF NOT EXISTS result(
                id INTEGER PRIMARY KEY AUTOINCREMENT,
                file_name TEXT NOT NULL,
                before_size INTEGER NOT NULL,
                after_size  INTEGER NOT NULL,
                start_time_unix INTEGER NOT NULL,
                finish_time_unix INTEGER NOT NULL
            );''')
            db.commit()
            db.execute('''INSERT INTO result (
                file_name, 
                before_size, 
                after_size, 
                start_time_unix, 
                finish_time_unix
            ) VALUES (?,?,?,?,?)''', (
                input_file,
                os.path.getsize(input_file),
                os.path.getsize(output_file),
                start_time,
                finish_time))
            db.commit()
        return result
    return wrapper


@record
def call_ffmpeg(input, output, arg_strs, debug=False):
    command = ['ffmpeg', '-i', input] + \
        arg_strs.split() + \
        ['-metadata', 'description={}@{}'.format(_name_, _version_), output]
    if debug:
        f = print
    else:
        f = call
    f(command)


def main():
    args = parse_args()
    if args.debug:
        print(args)

    for file_name, base_path in files(files=args.file, recursive=args.recursive):
        if not os.path.splitext(file_name)[1] in args.exts:
            continue
        if re.match(r'^.*\[.+_[^\]]+\].mp4$', file_name):
            continue

        full_input_file_name = os.path.join(base_path, file_name)

        desc = get_description_metadata(full_input_file_name)
        if desc and desc.startswith(_name_):
            continue

        output_name = os.path.splitext(file_name)[0] + '[x265_aac].mp4'
        full_output_file_name = os.path.join(base_path, output_name)

        call_ffmpeg(input=full_input_file_name,
                    output=full_output_file_name,
                    arg_strs=args.ffmpeg_args,
                    debug=args.debug)

        if not args.debug and args.delete:
            os.remove(full_input_file_name)


if __name__ == "__main__":
    main()
