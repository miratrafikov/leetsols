fn main() {
    let mut nums: [i32; 6] = [12,3,5,2,15,10];
    sort(&mut nums);
    for i in 0..nums.len() {
        println!("{}", nums[i]);
    }
}

fn sort(nums: &mut [i32]) {
    if nums.len() < 2 {
        return
    }

    for i in 0..nums.len() - 1 {
        let mut smallest = i;
        for j in i + 1..nums.len() {
            if nums[j] < nums[smallest] {
                smallest = j;
            }
        }
        
        let buff = nums[i];
        nums[i] = nums[smallest];
        nums[smallest] = buff;
    }
}
