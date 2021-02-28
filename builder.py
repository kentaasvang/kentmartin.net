

class Builder:
    data = ""

    @classmethod
    def add(cls, data):
        cls.data += data

    @classmethod
    def build(cls, filename="test.html"):
        with open(filename, "w") as file_handler:
            file_handler.write(cls.data)
