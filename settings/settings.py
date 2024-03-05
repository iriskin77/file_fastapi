import os
from dotenv import load_dotenv
from pathlib import Path

BASE_DIR = Path(__file__).resolve().parent.parent

print(BASE_DIR)
load_dotenv(os.path.join(BASE_DIR, ".env"))

absolute_path = os.path.abspath('file/apps')

PATH_FILE_STORAGE = os.path.join(os.path.dirname(os.path.dirname(__file__)), "apps/file/file_storage")


DATABASE_URI = f'postgres://{os.environ.get("POSTGRES_USER")}:' \
               f'{os.environ.get("POSTGRES_PASSWORD")}@' \
               f'{os.environ.get("POSTGRES_HOST")}:{os.environ.get("POSTGRES_PORT")}/' \
               f'{os.environ.get("POSTGRES_DB")}'


APPS_MODELS = [
    "apps.file.models",
    "apps.user.models",
    "aerich.models",
]

print(PATH_FILE_STORAGE)
#
# print(PATH_FILE_STORAGE)
# print(os.path.isdir(PATH_FILE_STORAGE))
# print(os.mkdir(PATH_FILE_STORAGE))