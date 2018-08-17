#include <stdio.h>

main(int argc, char const *argv[])
{
    int s[100],i,j,a,b,c;
	for (i=0;;i++)
	{
		scanf("%d",&s[i]);
		if (s[i]==0)
			break;
	}
	for (j=0;j<i;j++)
	{
		a=s[j]%10;
		b=s[j]/10%10;
		c=s[j]/100;
		if (a*a*a + b*b*b + c*c*c == s[j])
			printf("Yes\n");
		else
			printf("No\n");
	}
}

int dfs(int deep){
    
}
