let Array_Generic : Array<string> = ['1','3']


// 接口也可以用来描述数组

interface NumberArray {
    [index: number]: number;
}
let fibonacci: NumberArray = [1, 1, 2, 3, 5];
//NumberArray 表示：只要 index 的类型是 number，那么值的类型必须是 number

/**
 * -------------------------------------------------------------
 */

//一个比较常见的做法是，用 any 表示数组中允许出现任意类型：
let list: any[] = ['Ethan jay', 25, { website: 'http://github.com/7Ethan'}];