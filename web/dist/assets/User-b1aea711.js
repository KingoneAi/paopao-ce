import{_ as je,a as Ee}from"./post-item.vue_vue_type_style_index_0_lang-c8a377f9.js";import{_ as Ge}from"./post-skeleton-4fc3749c.js";import{D as Je,_ as ae,E as Ye,F as Ke,e as O,G as Qe,H as Xe,I as Ze,J as es}from"./index-d5021901.js";import{R as de,H as me,S as fe,b as he,e as te,i as ge,T as ss,F as as,a as ts,j as ie,o as ls,M as os,O as ns,k as us,f as is,g as cs,I as rs,G as _s}from"./naive-ui-d8de3dda.js";import{d as le,H as o,e as r,q as y,w as n,j as _,k as l,A as w,x as g,R as vs,c as ps,b as ds,E as ms,r as fs,f as k,Y as P,bf as h,y as ce,F as re,u as _e,h as ve}from"./@vue-a481fc63.js";import{_ as hs}from"./main-nav.vue_vue_type_style_index_0_lang-cc30e385.js";import{u as gs}from"./vuex-44de225f.js";import{b as ws}from"./vue-router-e5a2430e.js";import{b as ks}from"./formatTime-4210fcd1.js";import{W as ys}from"./v3-infinite-loading-2c58ec2f.js";import{i as bs,x as Ps,y as pe,z as $s,u as xs,D as Ts,G as zs}from"./@vicons-33f992ab.js";import"./content-ec46ab9c.js";import"./paopao-video-player-2fe58954.js";import"./copy-to-clipboard-4ef7d3eb.js";import"./@babel-725317a4.js";import"./toggle-selection-93f4ad84.js";import"./axios-4a70c6fc.js";/* empty css               */import"./seemly-76b7b838.js";import"./vueuc-39372edb.js";import"./evtd-b614532e.js";import"./@css-render-7124a1a5.js";import"./vooks-6d99783e.js";import"./vdirs-b0483831.js";import"./@juggle-41516555.js";import"./css-render-6a5c5852.js";import"./@emotion-8a8e73f6.js";import"./lodash-es-8412e618.js";import"./treemate-25c27bff.js";import"./async-validator-dee29e8b.js";import"./date-fns-975a2d8f.js";import"./moment-2ab8298d.js";const Us={class:"whisper-wrap"},Is={class:"whisper-line"},Fs={class:"whisper-line send-wrap"},Cs=le({__name:"whisper",props:{show:{type:Boolean,default:!1},user:{}},emits:["success"],setup(R,{emit:$}){const d=R,v=o(""),m=o(!1),u=()=>{$("success")},p=()=>{m.value=!0,Je({user_id:d.user.id,content:v.value}).then(s=>{window.$message.success("发送成功"),m.value=!1,v.value="",u()}).catch(s=>{m.value=!1})};return(s,i)=>{const x=de,T=me,a=fe,z=he,U=te,I=ge;return r(),y(I,{show:s.show,"onUpdate:show":u,class:"whisper-card",preset:"card",size:"small",title:"私信","mask-closable":!1,bordered:!1,style:{width:"360px"}},{default:n(()=>[_("div",Us,[l(a,{"show-icon":!1},{default:n(()=>[w(" 即将发送私信给: "),l(T,{style:{"max-width":"100%"}},{default:n(()=>[l(x,{type:"success"},{default:n(()=>[w(g(s.user.nickname)+"@"+g(s.user.username),1)]),_:1})]),_:1})]),_:1}),_("div",Is,[l(z,{type:"textarea",placeholder:"请输入私信内容（请勿发送不和谐内容，否则将会被封号）",autosize:{minRows:5,maxRows:10},value:v.value,"onUpdate:value":i[0]||(i[0]=F=>v.value=F),maxlength:"200","show-count":""},null,8,["value"])]),_("div",Fs,[l(U,{strong:"",secondary:"",type:"primary",loading:m.value,onClick:p},{default:n(()=>[w(" 发送 ")]),_:1},8,["loading"])])])]),_:1},8,["show"])}}});const Ms=ae(Cs,[["__scopeId","data-v-0cbfe47c"]]),qs={class:"whisper-wrap"},Ss={class:"whisper-line"},Ws={class:"whisper-line send-wrap"},Ls=le({__name:"whisper-add-friend",props:{show:{type:Boolean,default:!1},user:{}},emits:["success"],setup(R,{emit:$}){const d=R,v=o(""),m=o(!1),u=()=>{$("success")},p=()=>{m.value=!0,Ye({user_id:d.user.id,greetings:v.value}).then(s=>{window.$message.success("发送成功"),m.value=!1,v.value="",u()}).catch(s=>{m.value=!1})};return(s,i)=>{const x=de,T=me,a=fe,z=he,U=te,I=ge;return r(),y(I,{show:s.show,"onUpdate:show":u,class:"whisper-card",preset:"card",size:"small",title:"申请添加朋友","mask-closable":!1,bordered:!1,style:{width:"360px"}},{default:n(()=>[_("div",qs,[l(a,{"show-icon":!1},{default:n(()=>[w(" 发送添加朋友申请给: "),l(T,{style:{"max-width":"100%"}},{default:n(()=>[l(x,{type:"success"},{default:n(()=>[w(g(s.user.nickname)+"@"+g(s.user.username),1)]),_:1})]),_:1})]),_:1}),_("div",Ss,[l(z,{type:"textarea",placeholder:"请输入真挚的问候语",autosize:{minRows:5,maxRows:10},value:v.value,"onUpdate:value":i[0]||(i[0]=F=>v.value=F),maxlength:"120","show-count":""},null,8,["value"])]),_("div",Ws,[l(U,{strong:"",secondary:"",type:"primary",loading:m.value,onClick:p},{default:n(()=>[w(" 发送 ")]),_:1},8,["loading"])])])]),_:1},8,["show"])}}});const Os=ae(Ls,[["__scopeId","data-v-60be56a2"]]),Rs={key:0,class:"profile-baseinfo"},As={class:"avatar"},Bs={class:"base-info"},Ds={class:"username"},Hs={class:"userinfo"},Ns={class:"info-item"},Vs={class:"info-item"},js={class:"userinfo"},Es={class:"info-item"},Gs={class:"info-item"},Js={key:0,class:"user-opts"},Ys={key:0,class:"skeleton-wrap"},Ks={key:1},Qs={key:0,class:"empty-wrap"},Xs={key:1},Zs={key:2},ea={class:"load-more-wrap"},sa={class:"load-more-spinner"},aa=le({__name:"User",setup(R){const $=ss(),d=gs(),v=ws(),m="true".toLowerCase()==="true",u=o(!1),p=o(!1),s=vs({id:0,avatar:"",username:"",nickname:"",is_admin:!1,is_friend:!0,is_following:!1,created_on:0,follows:0,followings:0,status:1}),i=o(!1),x=o(!1),T=o(!1),a=o([]),z=o([]),U=o([]),I=o([]),F=o([]),B=o([]),C=o(v.query.s||""),t=o(+v.query.p||1),q=o("post"),D=o(+v.query.p||1),H=o(1),N=o(1),V=o(1),j=o(1),f=o(20),c=o(0),E=o(0),G=o(0),J=o(0),Y=o(0),K=o(0),we=()=>{p.value=!1,a.value=[],z.value=[],U.value=[],I.value=[],F.value=[],B.value=[],q.value="post",t.value=1,D.value=1,H.value=1,N.value=1,V.value=1,j.value=1,c.value=0,E.value=0,G.value=0,J.value=0,Y.value=0,K.value=0},ke=()=>{switch(q.value){case"post":A();break;case"comment":Q();break;case"highlight":X();break;case"media":Z();break;case"star":ee();break}},A=()=>{u.value=!0,O({username:C.value,style:"post",page:t.value,page_size:f.value}).then(e=>{u.value=!1,e.list.length===0&&(p.value=!0),t.value>1?a.value=a.value.concat(e.list):(a.value=e.list||[],window.scrollTo(0,0)),c.value=Math.ceil(e.pager.total_rows/f.value),z.value=a.value,E.value=c.value}).catch(e=>{a.value=[],t.value>1&&t.value--,u.value=!1})},Q=()=>{u.value=!0,O({username:C.value,style:"comment",page:t.value,page_size:f.value}).then(e=>{u.value=!1,e.list.length===0&&(p.value=!0),t.value>1?a.value=a.value.concat(e.list):(a.value=e.list||[],window.scrollTo(0,0)),c.value=Math.ceil(e.pager.total_rows/f.value),U.value=a.value,G.value=c.value}).catch(e=>{a.value=[],t.value>1&&t.value--,u.value=!1})},X=()=>{u.value=!0,O({username:C.value,style:"highlight",page:t.value,page_size:f.value}).then(e=>{u.value=!1,e.list.length===0&&(p.value=!0),t.value>1?a.value=a.value.concat(e.list):(a.value=e.list||[],window.scrollTo(0,0)),c.value=Math.ceil(e.pager.total_rows/f.value),I.value=a.value,J.value=c.value}).catch(e=>{a.value=[],t.value>1&&t.value--,u.value=!1})},Z=()=>{u.value=!0,O({username:C.value,style:"media",page:t.value,page_size:f.value}).then(e=>{u.value=!1,e.list.length===0&&(p.value=!0),t.value>1?a.value=a.value.concat(e.list):(a.value=e.list||[],window.scrollTo(0,0)),c.value=Math.ceil(e.pager.total_rows/f.value),F.value=a.value,Y.value=c.value}).catch(e=>{a.value=[],t.value>1&&t.value--,u.value=!1})},ee=()=>{u.value=!0,O({username:C.value,style:"star",page:t.value,page_size:f.value}).then(e=>{u.value=!1,e.list.length===0&&(p.value=!0),t.value>1?a.value=a.value.concat(e.list):(a.value=e.list||[],window.scrollTo(0,0)),c.value=Math.ceil(e.pager.total_rows/f.value),B.value=a.value,K.value=c.value}).catch(e=>{a.value=[],t.value>1&&t.value--,u.value=!1})},ye=e=>{switch(q.value=e,q.value){case"post":a.value=z.value,t.value=D.value,c.value=E.value,A();break;case"comment":a.value=U.value,t.value=H.value,c.value=G.value,Q();break;case"highlight":a.value=I.value,t.value=N.value,c.value=J.value,X();break;case"media":a.value=F.value,t.value=V.value,c.value=Y.value,Z();break;case"star":a.value=B.value,t.value=j.value,c.value=K.value,ee();break}},W=()=>{i.value=!0,Ke({username:C.value}).then(e=>{i.value=!1,s.id=e.id,s.avatar=e.avatar,s.username=e.username,s.nickname=e.nickname,s.is_admin=e.is_admin,s.is_friend=e.is_friend,s.created_on=e.created_on,s.is_following=e.is_following,s.follows=e.follows,s.followings=e.followings,s.status=e.status,ke()}).catch(e=>{i.value=!1,console.log(e)})},be=()=>{switch(q.value){case"post":D.value=t.value,A();break;case"comment":H.value=t.value,Q();break;case"highlight":N.value=t.value,X();break;case"media":V.value=t.value,Z();break;case"star":j.value=t.value,ee();break}},Pe=()=>{x.value=!0},$e=()=>{T.value=!0},xe=()=>{x.value=!1},Te=()=>{T.value=!1},M=e=>()=>ve(ie,null,{default:()=>ve(e)}),ze=ps(()=>{let e=[{label:"私信",key:"whisper",icon:M(Ps)}];return d.state.userInfo.is_admin&&(s.status===1?e.push({label:"禁言",key:"banned",icon:M(pe)}):e.push({label:"解封",key:"deblocking",icon:M(pe)})),s.is_following?e.push({label:"取消关注",key:"unfollow",icon:M($s)}):e.push({label:"关注",key:"follow",icon:M(xs)}),m&&(s.is_friend?e.push({label:"删除好友",key:"delete",icon:M(Ts)}):e.push({label:"添加朋友",key:"requesting",icon:M(zs)})),e}),Ue=e=>{switch(e){case"whisper":Pe();break;case"delete":Ie();break;case"requesting":$e();break;case"follow":case"unfollow":Fe();break;case"banned":case"deblocking":Ce();break}},Ie=()=>{$.warning({title:"删除好友",content:"将好友 “"+s.nickname+"” 删除，将同时删除 点赞/收藏 列表中关于该朋友的 “好友可见” 推文",positiveText:"确定",negativeText:"取消",onPositiveClick:()=>{i.value=!0,Qe({user_id:s.id}).then(e=>{i.value=!1,s.is_friend=!1,A()}).catch(e=>{i.value=!1,console.log(e)})}})},Fe=()=>{$.success({title:"提示",content:"确定"+(s.is_following?"取消关注":"关注")+"该用户吗？",positiveText:"确定",negativeText:"取消",onPositiveClick:()=>{i.value=!0,s.is_following?Xe({user_id:s.id}).then(e=>{i.value=!1,window.$message.success("取消关注成功"),W()}).catch(e=>{i.value=!1,console.log(e)}):Ze({user_id:s.id}).then(e=>{i.value=!1,window.$message.success("关注成功"),W()}).catch(e=>{i.value=!1,console.log(e)})}})},Ce=()=>{$.warning({title:"警告",content:"确定对该用户进行"+(s.status===1?"禁言":"解封")+"处理吗？",positiveText:"确定",negativeText:"取消",onPositiveClick:()=>{i.value=!0,es({id:s.id,status:s.status===1?2:1}).then(e=>{i.value=!1,s.status===1?window.$message.success("禁言成功"):window.$message.success("解封成功"),W()}).catch(e=>{i.value=!1,console.log(e)})}})},Me=()=>{t.value<c.value||c.value==0?(p.value=!1,t.value++,be()):p.value=!0};return ds(()=>{W()}),ms(()=>({path:v.path,query:v.query}),(e,b)=>{b.path==="/u"&&e.path==="/u"&&(C.value=v.query.s||"",we(),W())}),(e,b)=>{const qe=hs,Se=ls,se=os,oe=fs("router-link"),We=te,Le=ns,Oe=Ms,ne=us,L=is,Re=cs,Ae=Ge,Be=rs,De=je,ue=_s,He=Ee,Ne=as,Ve=ts;return r(),k("div",null,[l(qe,{title:"用户详情"}),l(Ne,{class:"main-content-wrap profile-wrap",bordered:""},{default:n(()=>[l(ne,{show:i.value},{default:n(()=>[s.id>0?(r(),k("div",Rs,[_("div",As,[l(Se,{size:72,src:s.avatar},null,8,["src"])]),_("div",Bs,[_("div",Ds,[_("strong",null,g(s.nickname),1),_("span",null," @"+g(s.username),1),m&&h(d).state.userInfo.id>0&&h(d).state.userInfo.username!=s.username&&s.is_friend?(r(),y(se,{key:0,class:"top-tag",type:"info",size:"small",round:""},{default:n(()=>[w(" 好友 ")]),_:1})):P("",!0),h(d).state.userInfo.id>0&&h(d).state.userInfo.username!=s.username&&s.is_following?(r(),y(se,{key:1,class:"top-tag",type:"success",size:"small",round:""},{default:n(()=>[w(" 已关注 ")]),_:1})):P("",!0),s.is_admin?(r(),y(se,{key:2,class:"top-tag",type:"error",size:"small",round:""},{default:n(()=>[w(" 管理员 ")]),_:1})):P("",!0)]),_("div",Hs,[_("span",Ns,"UID. "+g(s.id),1),_("span",Vs,g(h(ks)(s.created_on))+" 加入",1)]),_("div",js,[_("span",Es,[l(oe,{onClick:b[0]||(b[0]=ce(()=>{},["stop"])),class:"following-link",to:{name:"following",query:{s:s.username,n:s.nickname,t:"follows"}}},{default:n(()=>[w(" 关注  "+g(s.follows),1)]),_:1},8,["to"])]),_("span",Gs,[l(oe,{onClick:b[1]||(b[1]=ce(()=>{},["stop"])),class:"following-link",to:{name:"following",query:{s:s.username,n:s.nickname,t:"followings"}}},{default:n(()=>[w(" 粉丝  "+g(s.followings),1)]),_:1},8,["to"])])])]),h(d).state.userInfo.id>0&&h(d).state.userInfo.username!=s.username?(r(),k("div",Js,[l(Le,{placement:"bottom-end",trigger:"click",size:"small",options:ze.value,onSelect:Ue},{default:n(()=>[l(We,{quaternary:"",circle:""},{icon:n(()=>[l(h(ie),null,{default:n(()=>[l(h(bs))]),_:1})]),_:1})]),_:1},8,["options"])])):P("",!0)])):P("",!0),l(Oe,{show:x.value,user:s,onSuccess:xe},null,8,["show","user"]),l(Os,{show:T.value,user:s,onSuccess:Te},null,8,["show","user"])]),_:1},8,["show"]),l(Re,{class:"profile-tabs-wrap",type:"line",animated:"",value:q.value,"onUpdate:value":ye},{default:n(()=>[l(L,{name:"post",tab:"泡泡"}),l(L,{name:"comment",tab:"评论"}),l(L,{name:"highlight",tab:"亮点"}),l(L,{name:"media",tab:"图文"}),l(L,{name:"star",tab:"喜欢"})]),_:1},8,["value"]),u.value&&a.value.length===0?(r(),k("div",Ys,[l(Ae,{num:f.value},null,8,["num"])])):(r(),k("div",Ks,[a.value.length===0?(r(),k("div",Qs,[l(Be,{size:"large",description:"暂无数据"})])):P("",!0),h(d).state.desktopModelShow?(r(),k("div",Xs,[(r(!0),k(re,null,_e(a.value,S=>(r(),y(ue,{key:S.id},{default:n(()=>[l(De,{post:S},null,8,["post"])]),_:2},1024))),128))])):(r(),k("div",Zs,[(r(!0),k(re,null,_e(a.value,S=>(r(),y(ue,{key:S.id},{default:n(()=>[l(He,{post:S},null,8,["post"])]),_:2},1024))),128))]))]))]),_:1}),c.value>0?(r(),y(Ve,{key:0,justify:"center"},{default:n(()=>[l(h(ys),{class:"load-more",slots:{complete:"没有更多泡泡了",error:"加载出错"},onInfinite:b[2]||(b[2]=S=>Me())},{spinner:n(()=>[_("div",ea,[p.value?P("",!0):(r(),y(ne,{key:0,size:14})),_("span",sa,g(p.value?"没有更多泡泡了":"加载更多"),1)])]),_:1})]),_:1})):P("",!0)])}}});const Wa=ae(aa,[["__scopeId","data-v-0436abd3"]]);export{Wa as default};
