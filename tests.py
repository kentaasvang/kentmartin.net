import os
from builder import Builder

TESTFILE="TESTFILE"

def test_builder():
    Builder.add("first ")
    Builder.add("second ")
    Builder.add("third ")

    # produces a file that needs to be cleaned up
    Builder.build(TESTFILE)

    expected = "first second third "
    
    with open(TESTFILE, "r") as file_handler:
        result = file_handler.read()

    assert expected == result, "build() did not produce correct file-content"

    # clean up file
    os.remove(TESTFILE)


if __name__ == "__main__":
    test_builder()