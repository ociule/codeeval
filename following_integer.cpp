/**
*/

#include <iostream>
#include <fstream>
#include <sstream>

std::vector<int> * getDigits(int n) {

}

int getNextLargest(int n) {
    return n + 1;
}

int main(int argc, char *argv[])
{
    std::ifstream file;
    std::string lineBuffer;

    if (argc < 2) {
        std::cout << "Please add the name of the file to process." << std::endl;
        exit(1);
    }

    file.open(argv[1], std::ifstream::in);

    while (!file.eof()) 
    {
            getline(file, lineBuffer);
            if (lineBuffer.length() == 0)
                continue;
            else 
            {

                std::istringstream iss(lineBuffer);
                int val;
                iss >> val;
                std::cout << val << " " << getNextLargest(val) << std::endl;
            }
    }
    return 0;
}
