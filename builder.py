

class Builder:
    data = ""

    @classmethod
    def add(cls, data):
        cls.data += data

    @classmethod
    def build(cls, filename="test.html"):
        with open(filename, "w") as file_handler:
            file_handler.write(cls.data)


if __name__ == "__main__":
    Builder.add("<header></header>")
    Builder.add("<body></body>")
    Builder.add("<footer></footer>")

    Builder.build()