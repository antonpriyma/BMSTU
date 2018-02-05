unsigned long peak(unsigned long nel, int (*less)(unsigned long i, unsigned long j))
{
    unsigned long midle,start=0,finish=nel-1;

    

    while (finish>start)
    {
        midle = start/2 + finish/2;
        if (less(midle,midle+1))
            start = midle+1;
        else
            if (less(midle,midle-1))
                finish = midle -1;
                else
                 return midle;        
    }
    //Что будет, если убрать return mid? 
    return finish;


}
