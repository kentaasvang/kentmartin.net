from builder import Builder
from website import Index


def create_page():

    index = Index()

    for i in ["header", "main", "footer"]:
        with open(i, "r") as file_handler:
            Builder.add(file_handler.read())

    Builder.build(index.name)


if __name__ == "__main__":
    create_page()