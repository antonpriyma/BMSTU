
int maxarray(void *base, unsigned long nel, unsigned long width,
        int (*compare)(void *a, void *b))
        {
            int j,i;
            void *max;
            max = base;//����� �������� ���������
            j=0;
            for (i=0; i<(nel)*width;i+=width)
                if ((compare(max,((char*)base+i)))<0)
                {
                    max = ((char*)base) + i;//������������ ������ ��������� ������ �������� ��������
                    j=i/width;
                }

                return j;
        }
