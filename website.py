

class BaseWebsite:
    def __init__(self, name=None):
        self.name=name


class Index(BaseWebsite):
    def __init__(self):
        super(Index, self).__init__("index.html")

