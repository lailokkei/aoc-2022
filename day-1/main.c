#include <string.h>
#include <stdlib.h>
#include <stdio.h>
const int size = 256;
int main() {
    FILE * fp;
    
    fp = fopen("calories.txt", "r");
    if (fp == NULL) {
        perror("failed");
        return 1;
    }
    char buffer[size];

    int max = 0;
    int current = 0;
    while(fgets(buffer, size, fp)){

        buffer[strcspn(buffer, "\n")] = '\0';

        if (strcmp(buffer, "") == 0) {
            if (current > max) {
                max = current;
            }
            current = 0;
        }
        
        current += atoi(buffer);
    }

    printf("%d\n", max);
}
