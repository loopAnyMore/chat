package filter

import (
	"regexp"
)


//中间屏蔽长度
const  fifterContain int=10
//单例
var m *Chatfilter
	
func GetInstance() *Chatfilter{
	if m==nil{
		m=&Chatfilter{}
		reg,_:=regexp.Compile("[^\u4e00-\u9fa5]")
		m.reg=reg
	}
	return m
}

type Chatfilter struct{
	reg *regexp.Regexp
	root *node
}


func(f *Chatfilter) Deal(txt string,replaceChar rune)string{
	 key:=[]rune(txt)
	 length:=len(key)
	 for i:=0;i<length;i++{
		 index,checkNum:=f.getReplaceIndex(key,length,i)
		 if checkNum>0{
			for b:=i;b<=index;b++{
				key[b]=replaceChar
			}
			i=index-1
		 }
	 }
	return  string(key)
}

func(f *Chatfilter) getReplaceIndex(key []rune ,length int,beginIndex int) (int,int){
	matchNum:=0
	ignore:=false
	ignoreNum:=0
	var ignoreRoot *node
	// 敏感词结束标识位：用于敏感词只有1位的情况
	flag:=false
	nowRoot:=f.root 
	var endIndex int
	for i:=beginIndex;i<length;i++{
		endIndex=i
		nowKey:=key[i]
		if val,ok:=nowRoot.child[nowKey];ok{
				nowRoot=val
			if	ignoreRoot==nil{
				ignoreRoot=nowRoot
				ignore=true
			}	
			//此时判断是否为子集关系
			if(ignoreNum>0){
				if _,ok2:=ignoreRoot.child[nowKey];ok2{
						matchNum++
				}
			}
			ignoreNum=0;
			matchNum++
			if nowRoot.end==true {
				// 结束标志位为true
				flag = true
				break
			}
		ignoreRoot=nowRoot
		}else{
			if ignore{
				if(ignoreNum<fifterContain){
					// //非中文
					// if f.reg.MatchString(word2){
						ignoreNum++
						continue
					//}
				}
				break
			}
	   		break
		}
	}
	if !flag{
		matchNum=0
	}
	return endIndex,matchNum;
}

type node struct{
	child map[rune] *node
	end bool	
}

func (f *Chatfilter) Insert(txt string){
	if(len(txt)<1){
		return
	}
	if(f.root==nil){
		f.root=&node {
			child: make(map[rune]*node),
			end: false,
		}
	}
	now:=f.root
	for _,v:=range txt{
		if _,ok:=now.child[v];!ok{
			now.child[v]=&node {
				child: make(map[rune]*node),
				end: false,
			}
		}
		
		now=now.child[v]
	}
	now.end=true
}

