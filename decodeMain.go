package main

import (
	pb "gitlab.mobvista.com/decodeCreativePb/decodePb"
	"fmt"
	"github.com/gogo/protobuf/proto"
)

func main(){
	var creative,creative1,creative2 pb.Creative
	creative.Url = "com.cn"
	creative.Width = 10
	encode,_ := proto.Marshal(&creative)
	fmt.Println(encode)
	proto.Unmarshal(encode,&creative1)
	fmt.Println("creatve1=",creative1)

	str := "\nYhttp://cdn-adn.rayjump.com/cdn-adn/vh/18/08/15/20/38/9c488d977937575b35516ef6be3e8ae8.mp4\x10\x1e\x18\xc0\xcc\xb3\x02\"\t1920x1080(\x80\x0f0\xb8\b8d@\xbd\nj\tvideo/mp4r\a3549946x\xd3\x8e\xb1\xc2\t\x82\x01 6e355e8bd4a6fe34d9c9af1390c02a14\x88\x01\x02\xa2\x010sniper3d_v_201803_89_lang-ZHLL_dim-1920x1080.mp4"
	proto.Unmarshal([]byte(str),&creative2)
	fmt.Println("creatve2=",creative2)

}
