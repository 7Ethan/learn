import java.util.Arrays;

public class Hello {
    public void hello(){
        System.out.println("Hello Java");
    }

    public int[] getArray(int length){
        int[] arr = new int[length];
        for (int i = 0;i < arr.length;i++){
            arr[i] = (int)(Math.random() * 100);
        }
        return arr;
    }

    static int num = 0; //must not in local method
    // public  int staticVal(){
    //     static int num = 0;
    //     num++;
    //     return num;
    // }

    public class Inner{
        public void show(){
            System.out.println("Inner Class!");
        }
    }

    String name; // 声明变量name
	String sex; // 声明变量sex
	static int age;// 声明静态变量age
    
    // 构造方法
    // 构造方法
	public Hello() { 
		System.out.println("通过构造方法初始化name=========" );
		// name = "tom";
    }
    
	public Hello(String name) { 
		System.out.println("通过构造方法初始化name : " + name);
		// name = "tom";
	}
    
    // 初始化块
	{ 
		System.out.println("通过初始化块初始化sex");
		sex = "男";
	}
    
    // 静态初始化块
    /**
     * 静态初始化块只在类加载时执行，且只会执行一次，
     * 同时静态初始化块只能给静态变量赋值，不能初始化普通的成员变量。
     */
	static        { 
		System.out.println("通过静态初始化块初始化age");
		age = 20;
	}
    
	public void show() {
		System.out.println("姓名：" + name + "，性别：" + sex + "，年龄：" + age);
	}

    public static void main(String[] args) {
        Hello instance= new Hello();
        // instance.hello();
        // int arr[] = { 78, 93, 97, 84, 63 };
        // System.out.println(arr[0]);
        // int array[] = instance.getArray(10);
        // for(int num:array){
        //     System.out.println(num);
        // }

        // System.out.println(instance.staticVal());
        // System.out.println(instance.staticVal());
        // System.out.println(instance.staticVal());
        Inner i = instance.new Inner();
        i.show();

        // ----------------------
        HelloWorld h = new HelloWorld("sergey");
        System.out.println(h.toString());
    }

}

class HelloWorld extends Hello{
    public HelloWorld(String name){
        super();
        System.out.println("HELLOWORLD" + name);
    }

    public void info(){
        System.out.println("HelloWorld extends Hello");
    }
}