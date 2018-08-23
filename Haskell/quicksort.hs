import System.Random
-- import Data.Time.Clock as Clock

arr = [345,567,-56,90,34,9006,456,6,324,68,86,234]

qsort [] = [] --如果不指定将会运行出错: Non-exhaustive patterns......(无限递归)

qsort (first:rest) = 
    qsort (filter (< first) rest) ++ [first] ++ qsort (filter (>= first) rest)
    -- (first:rest) -->把传递的数组结构分成首项+剩余两个部分 
    -- 同时作为参数传递到后面的函数中，首项的作用就是上面说的first啦，当然也可以时最后一项，取首项时方便一些

    -- [a|a<-xs,a<=x] 是什么意思呢？ 
    -- a|，这个a代表输出，可以理解为返回值，a<-xs,a<=x表示这个a从xs中来，同时a小于x，
    -- 也就是传进来的去掉第一项的数组，输出小于等于x的数组元素，综合来看，这句话实现了：把小的数组元素扔左边的功能；

    

main = do
    -- start < Clock.getCurrentTime
    let newArray = qsort arr
    -- end < Clock.getCurrentTime
    print newArray
    -- let time = diffClockTime end start
    -- let time = Clock.secondsToDiffTime end start
    -- putStrLn("Cost time is :" ++ show(time))


    