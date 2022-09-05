#include<stdio.h>

void main()
{
    char *binaryPath = "/bin/echo";
    char *arg1 = "123";
    execl(binaryPath, binaryPath, arg1, NULL);
    printf("exec error\n");
}