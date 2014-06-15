#include <iostream>
#include <sys/stat.h>


long GetFileSize(std::string filename)
{
    struct stat stat_buf;
    int rc = stat(filename.c_str(), &stat_buf);
    return rc == 0 ? stat_buf.st_size : -1;
}

int main(int argc, char *argv[])
{
    std::cout << GetFileSize(argv[1]) << std::endl;
    return 0;
}
