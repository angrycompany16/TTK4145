// Compile with `gcc foo.c -Wall -std=gnu99 -lpthread`, or use the makefile
// The executable will be named `foo` if you use the makefile, or `a.out` if you use gcc directly

#include <pthread.h>
#include <stdio.h>
#include <semaphore.h>

int i = 0;

// Note the return type: void*
void* incrementingThreadFunction(){
    for (int n = 0; n < 1000000; n++) {
        i++;
    }
    return NULL;
}

void* decrementingThreadFunction(){
    for (int n = 0; n < 1000000; n++) {
        i--;
    }
    return NULL;
}


int main(){
    pthread_t inc_thread_id;
    pthread_create(&inc_thread_id, NULL, incrementingThreadFunction, NULL);
    pthread_t dec_thread_id;
    pthread_create(&dec_thread_id, NULL, decrementingThreadFunction, NULL);

    pthread_join(inc_thread_id, NULL);
    pthread_join(dec_thread_id, NULL);

    printf("The magic number is: %d\n", i);
    return 0;
}