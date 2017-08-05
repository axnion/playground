fn main() {
    let int_vec = vec![6,2,8,1,9,5,3];
    let new_int_vec = merge_sort(int_vec);

    // for element in new_int_vec {
    //     println!("{:?}", element);
    // }
}

fn merge_sort(mut int_vec: Vec<i32>) -> Vec<i32> {
    let length = int_vec.len();

    let mut before = String::new();

    for e in &int_vec {
        before.push_str(&" ".to_string());
        before.push_str(&e.to_string());
    }

    if length <= 1 {
        return int_vec;
    }

    let left = merge_sort(int_vec[0..length / 2].to_vec());
    let right = merge_sort(int_vec[length / 2..length].to_vec());

    let mut i_l = 0;
    let mut i_r = 0;

    int_vec.clear();

    loop {
        if i_l == left.len() {
            int_vec.push(right[i_r]);
            i_r = i_r + 1;
        } else if i_r == right.len() {
            int_vec.push(left[i_l]);
            i_l = i_l + 1;
        } else {
            if left <= right {
                int_vec.push(left[i_l]);
                i_l = i_l + 1;
            } else {
                int_vec.push(right[i_r]);
                i_r = i_r + 1;
            }
        }

        if i_l == left.len() && i_r == right.len() {
            break;
        }
    }

    println!("{:?} => {:?}", before, int_vec);

    return int_vec;
}

// fn merge_sort(int_array: &[i32]) {
//     let mut i = 0;
//     let length = int_array.len();
//
//     loop {
//         if i < length {
//             println!("{}: {}", i, int_array[i]);
//             i = i + 1;
//         } else {
//             break;
//         }
//     }
// }
