unsigned long binsearch(unsigned long nel, int (*compare)(unsigned long i))
{
    unsigned long start = 0;
    unsigned long finish = nel-1;
    unsigned long mid;
    
    mid = (start+finish)/2;
    
    while (compare(mid) &&  (finish-start>1)) {
        mid = (start+finish)/2;
        if (compare(mid)==-1)
            start = mid;
        if (compare(mid)==1)
            finish = mid;
    }
    
    if (!(compare(mid))) return mid;
    return nel;

}