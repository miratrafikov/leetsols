function hammingDistance(x: number, y: number): number {
    let mask = 1;
    let count = 0;
    for (let i = 0; i < 31; i++) {
        if ((x & mask) != (y & mask)) count++;
        mask = mask << 1;
    }
    return count;
};
