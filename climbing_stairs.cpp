/**
    This is NOT a correct solution! The output from countClimbingStairs overflows if the input n is large enough.
    To solve this, we would need a bigint library. Otherwise, the solution is a correct memoized fibonacci, FWIW.
*/

#include <iostream>
#include <fstream>
#include <sstream>
#include <map>

std::map<int, unsigned long long int> cache;


unsigned long long int countClimbingStairs(int n) {
    // This is the same as fibonacci(n), actually
    // We should do it in closed form
    if (n == 0) {
        return 0;
    }
    if (n == 1) {
        return 1;
    }
    if (n == 2) {
        return 2;
    }

    
    if (cache[n] > 0) {
        return cache[n];
    } else {
        int res = countClimbingStairs(n - 2) + countClimbingStairs(n - 1);
        cache[n] = res;
        return res;
    }
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

    for (int j = 0; j < 100; j++) {
        if (cache[j] > 0) {
            std::cout << j << " is not zeroed" << std::endl;
        }
    }

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
                std::cout << val << " " << countClimbingStairs(val) << std::endl;
            }
    }
    return 0;
}
