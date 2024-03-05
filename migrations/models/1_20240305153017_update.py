from tortoise import BaseDBAsyncClient


async def upgrade(db: BaseDBAsyncClient) -> str:
    return """
        ALTER TABLE "filetoupload" ALTER COLUMN "processed_at" DROP NOT NULL;"""


async def downgrade(db: BaseDBAsyncClient) -> str:
    return """
        ALTER TABLE "filetoupload" ALTER COLUMN "processed_at" SET NOT NULL;"""