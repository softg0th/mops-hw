from dataclasses import dataclass

import pymongo


@dataclass
class Config:
    url: str
    db: str
    column: str


def clear_database(db_configuration):
    client = pymongo.MongoClient(db_configuration.url)
    db = client[db_configuration.db]
    col = db[db_configuration.column]

    col.delete_many({'_id': {"$ne": None}})


if __name__ == "__main__":
    config = Config('mongodb://localhost:27017/', 'iot', 'messages')
    clear_database(config)
