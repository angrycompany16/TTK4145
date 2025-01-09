// Compile with `gcc foo.c -Wall -std=gnu99 -lpthread`, or use the makefile
// The executable will be named `foo` if you use the makefile, or `a.out` if you use gcc directly

#include <pthread.h>
#include <stdio.h>
#include <semaphore.h>

int i = 0;

// I changed my answer to using a mutex instead of a semaphore. The reason for this is
// that the thread that locks the mutex will also have to be the one that unlocks it, 
// there will not be any case where the resource should be unlocked from another
// thread.
pthread_mutex_t lock;

// Note the return type: void*
void* incrementingThreadFunction(){
    for (int n = 0; n < 1000000; n++) {
        pthread_mutex_lock(&lock);
        i++;
        pthread_mutex_unlock(&lock);
    }
    return NULL;
}

void* decrementingThreadFunction(){
    for (int n = 0; n < 1000000; n++) {
        pthread_mutex_lock(&lock);
        i--;
        pthread_mutex_unlock(&lock);
    }
    return NULL;
}


int main(){
    pthread_t inc_thread_id;
    pthread_create(&inc_thread_id, NULL, incrementingThreadFunction, NULL);
    pthread_t dec_thread_id;
    pthread_create(&dec_thread_id, NULL, decrementingThreadFunction, NULL);

    pthread_mutex_init(&lock, 0);

    pthread_join(inc_thread_id, NULL);
    pthread_join(dec_thread_id, NULL);

    pthread_mutex_destroy(&lock);
    printf("The magic number is: %d\n", i);
    return 0;
}
