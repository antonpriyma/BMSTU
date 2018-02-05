#include <string.h>
#include <malloc.h>

void revarray(void* base, unsigned long nel, unsigned long width)
{
        char* memoryBlock = (char*)malloc(nel*width); //пам€ть дл€ перегонки
        //обход массива с конца
        int i,j; //i - позици€ в base, i - позици€ в memoryBlock
        for(i = (nel-1)*width,j = 0; i > -1; i = i-width, j = j+width)
        {
                memcpy(memoryBlock+j, ((char*)base)+i, width); //скопировать очередной элемент
        }
        memcpy(base, memoryBlock, nel*width);
        free(memoryBlock);
}