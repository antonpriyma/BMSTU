#include <string.h>
#include <malloc.h>

void revarray(void* base, unsigned long nel, unsigned long width)
{
        char* memoryBlock = (char*)malloc(nel*width); //������ ��� ���������
        //����� ������� � �����
        int i,j; //i - ������� � base, i - ������� � memoryBlock
        for(i = (nel-1)*width,j = 0; i > -1; i = i-width, j = j+width)
        {
                memcpy(memoryBlock+j, ((char*)base)+i, width); //����������� ��������� �������
        }
        memcpy(base, memoryBlock, nel*width);
        free(memoryBlock);
}