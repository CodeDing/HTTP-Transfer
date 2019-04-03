
###
   HTTP Transfer
###   


####
  Request
####

```
   curl -v "http://localhost:8000/api/v1/selector?ad_network=bdm5&data=%7B%22app%22%3A%7B%22package_name%22%3A%22%22%2C%22version%22%3A%221.0.0%22%7D%2C%22selector%22%3A%7B%22package_version%22%3A%221.0.0%22%2C%22app_id%22%3A%22c0bd2a44%22%2C%22ad_id%22%3A1413602%2C%22app_name%22%3A%22%E8%87%AA%E8%B4%A1%E8%AE%BA%E5%9D%9B%22%2C%22app_storeurl%22%3Anull%2C%22slot_id%22%3A%226041732%22%2C%22slot_type%22%3A3%2C%22package_name%22%3A%22com.mocuz.zigongluntan%22%7D%2C%22id%22%3A%225000734.1-23.1po3dom.3.2.ppdiqb.bf31%22%2C%22ext%22%3A%7B%22filter_curl%22%3A%22%22%7D%2C%22slot%22%3A%7B%22current_selector%22%3A%22bdm5%22%2C%22id%22%3A%225000734%22%2C%22width%22%3A960%2C%22media_type%22%3A%22app%22%2C%22height%22%3A640%2C%22https%22%3A0%7D%2C%22user%22%3A%7B%22id%22%3A%22fc8fd90d77ba3cfe9c9074717f651261%22%7D%2C%22device%22%3A%7B%22operator%22%3A1%2C%22os%22%3A%22android%22%2C%22physical_width%22%3A1920%2C%22mac%22%3A%2200%3A00%3A00%3A00%3A00%3A00%22%2C%22model%22%3A%22Nexus%206P%22%2C%22location%22%3A%7B%22timestamp%22%3A1554275603611%2C%22lng%22%3A%220.0%22%2C%22lat%22%3A%220.0%22%7D%2C%22ip%22%3A%22222.66.148.6%22%2C%22os_version%22%3A%226.0%22%2C%22oreintation%22%3A0%2C%22ua%22%3A%22Mozilla%2F5.0%20%28Linux%3B%20Android%205.1%3B%20OPPO%20R9tm%20Build%2FLMY47I%3B%20wv%29%20AppleWebKit%2F537.36%20%28KHTML%2C%20like%20Gecko%29%20Version%2F4.0%20Chrome%2F43.0.2357.121%20Mobile%20Safari%2F537.36%22%2C%22type%22%3A1%2C%22imei%22%3A%22864690024387832%22%2C%22vendor%22%3A%22Huawei%22%2C%22physical_height%22%3A1080%2C%22physical_density%22%3A%221.0%22%2C%22network%22%3A1%2C%22android_id%22%3A%225c9fb234d9ef1d67%22%7D%2C%22region%22%3A%7B%22country%22%3A%22%E4%B8%AD%E5%9B%BD%22%2C%22province%22%3A%22%E4%B8%8A%E6%B5%B7%22%2C%22city%22%3A%22%E4%B8%8A%E6%B5%B7%E5%B8%82%22%7D%7D"
```

####
  Response
####

```
  {
      "code": 0,
      "msg": "",
      "id": "5000734.1-94.1po3dom.3.2.ppdhli.c4b8",
      "width": 960,
      "height": 640,
      "ad_id": 1413602,
      "slot_id": "5000734",
      "content": {
          "type": "text-icon",
          "action": 1,
          "url": "http://f10.baidu.com/it/u=329052672,1992901844&fm=76",
          "c_url": "http://cpro.baidu.com/cpro/ui/uijs.php?en=mywWUA71T1Yknzu9TZKxpyfqn1c1n1mznWR4PBu9TZKxTv-b5ynsmhfzmHfYFhPC5H0huAbqnH0YrjRLPBuGTdq9TZ0qniuJp1dbnAubuHRYuyc1uWRkryNxnauo5iNanzN7PzNjnzNAPzNjnBNjfzNaPiNDraNanzNKPaNaPiNafiuonjY-fWn-wHT-f1n-wWT-f1c-fYn-fWR-wj6-fWn-fHf-fWR-fbDhUZNopHYkFhdWTAYqn1R1nWDhTHdWnAFbnhDYP7qWTZchThcqniuzT1YkFMPbpdqv5HfhTvwogLu-TMPGUv3qPi3vQW0hTvN_UANzgv-b5HchTv-b5H--PHD1uhNBPHwbuHKhnAfhTLwGujYvnj0kFMfqIZKWUA-WpvNbndqVUvFxmgPsg1nhIAd15HDdPHfzP1fkn1fhIZRqPW0YnHT1nBud5y9YIZ0-nYD-nbm-nbmYmhwBPAD-nbf-nbwWnAFbnhDYPaRzwyPEUiuv5HDhpHdWP16krj63Ps&adx=1&c=news&cf=1&cp=chuilei_middle_page_app&ds=cmp&eid_list=1082_1085_1039&expid=6282_6562_9012_9072_9089_9503_9506_9511_9522_9560_9568_9580_9598_9599_12205_12210&fv=0&imei=864690024387832&imei_cp=0&img_typ=518&itm=0&k_words=%C2%CC%B5%D8%B3%A4%B5%BA%5F%C2%CC%B5%D8%B3%A4%B5%BA%5F%C2%CC%B5%D8%B3%A4%B5%BA%5F%C2%CC%B5%D8%B3%A4%B5%BA%5F%C2%CC%B5%D8%B3%A4%B5%BA%5F%C2%CC%B5%D8%B3%A4%B5%BA&kdi0=1048576&lu_idc=nj02&lukid=1&lus=9e513feb54de0f0d&lust=5ca45756&mscf=0&mtids=28706944,5185226,2000000013&n=10&nttp=3&os=android&p=baidu&sh=640&ssp2=0&sw=960&swi=1&tsf=dtp:2&u=&uicf=lurecv&urlid=0&vn=6&wi=4",
          "imp": [
              "http://wn.pos.baidu.com/adx.php?c=d25pZD05ZTUxM2ZlYjU0ZGUwZjBkXzAAcz05ZTUxM2ZlYjU0ZGUwZjBkAHQ9MTU1NDI3NDEzNABzZT0yAGJ1PTQAcHJpY2U9WEtSWFZnQUlacnA3akVwZ1c1SUE4dkU3djdac2NhMDhEWFpKdFEAY2hhcmdlX3ByaWNlPTIyMTYAc2hhcmluZ19wcmljZT0yMjE2MDAwAHdpbl9kc3A9NABjaG1kPTEAYmRpZD0AY3Byb2lkPQB3ZD0wAHR1PTYwNDE3MzIAYWRjbGFzcz0xMwBzcmN0PTQAcG9zPTAAYmNobWQ9MAB2PTEAaT1mMjYyOTIwMw"
          ],
          "clk": [
              "http://cpro.baidu.com/cpro/ui/uijs.php?en=mywWUA71T1Yknzu9TZKxpyfqn1c1n1mznWR4PBu9TZKxTv-b5ynsmhfzmHfYFhPC5H0huAbqnH0YrjRLPBuGTdq9TZ0qniuJp1dbnAubuHRYuyc1uWRkryNxnauo5iNanzN7PzNjnzNAPzNjnBNjfzNaPiNDraNanzNKPaNaPiNafiuonjY-fWn-wHT-f1n-wWT-f1c-fYn-fWR-wj6-fWn-fHf-fWR-fbDhUZNopHYkFhdWTAYqn1R1nWDhTHdWnAFbnhDYP7qWTZchThcqniuzT1YkFMPbpdqv5HfhTvwogLu-TMPGUv3qPi3vQW0hTvN_UANzgv-b5HchTv-b5H--PHD1uhNBPHwbuHKhnAfhTLwGujYvnj0kFMfqIZKWUA-WpvNbndqVUvFxmgPsg1nhIAd15HDdPHfzP1fkn1fhIZRqPW0YnHT1nBud5y9YIZ0-nYD-nbm-nbmYmhwBPAD-nbf-nbwWnAFbnhDYPaRzwyPEUiuv5HDhpHdWP16krj63Ps&adx=1&c=news&cf=1&cp=chuilei_middle_page_app&ds=cmp&eid_list=1082_1085_1039&expid=6282_6562_9012_9072_9089_9503_9506_9511_9522_9560_9568_9580_9598_9599_12205_12210&fv=0&imei=864690024387832&imei_cp=0&img_typ=518&itm=0&k_words=%C2%CC%B5%D8%B3%A4%B5%BA%5F%C2%CC%B5%D8%B3%A4%B5%BA%5F%C2%CC%B5%D8%B3%A4%B5%BA%5F%C2%CC%B5%D8%B3%A4%B5%BA%5F%C2%CC%B5%D8%B3%A4%B5%BA%5F%C2%CC%B5%D8%B3%A4%B5%BA&kdi0=1048576&lu_idc=nj02&lukid=1&lus=9e513feb54de0f0d&lust=5ca45756&mscf=0&mtids=28706944,5185226,2000000013&n=10&nttp=3&os=android&p=baidu&sh=640&ssp2=0&sw=960&swi=1&tsf=dtp:2&u=&uicf=lurecv&urlid=0&vn=6&wi=4"
          ],
          "title": "上海绿地长岛详情，其他人都不知道！",
          "desc": "上海绿地长岛详情",
          "ext_urls": [
              "http://t12.baidu.com/it/u=3843096377,2204229908&fm=76",
              "http://f12.baidu.com/it/u=2500509893,13256531&fm=76"
          ],
          "icon_url": "http://f10.baidu.com/it/u=329052672,1992901844&fm=76",
          "app": {
              "icon_url": "http://f10.baidu.com/it/u=329052672,1992901844&fm=76",
              "size": 438
          },
          "duration": 0
      },
      "union": "bdm5"
  }
```
