import{w as F,x as z,y as I,z as j,_ as E}from"./index-8275f3fe.js";import{p as U}from"./@vicons-4a4bea9c.js";import{d as $,r as i,n as q,j as A,a1 as x,o as c,c as _,L as n,$ as s,K as b,e as L,M as f,O as u,Y as h,w as D,a6 as K,F as Y,a2 as G}from"./@vue-0075cdd2.js";import{o as H,M as B,j as J,e as P,O as Q,L as R,F as W,f as X,g as Z,a as ee,k as oe}from"./naive-ui-18c654e5.js";import{_ as te}from"./main-nav.vue_vue_type_style_index_0_lang-899ed46d.js";import{u as ne}from"./vuex-9af76a69.js";import"./vue-router-863a6505.js";import"./axios-4a70c6fc.js";/* empty css               */import"./seemly-76b7b838.js";import"./vueuc-471a8e21.js";import"./evtd-b614532e.js";import"./@css-render-d7709484.js";import"./vooks-848f90ef.js";import"./vdirs-b0483831.js";import"./@juggle-41516555.js";import"./css-render-6a5c5852.js";import"./@emotion-8a8e73f6.js";import"./lodash-es-8412e618.js";import"./treemate-25c27bff.js";import"./async-validator-dee29e8b.js";import"./date-fns-975a2d8f.js";const se={key:0,class:"tag-item"},ae={key:0,class:"tag-quote"},ce={key:1,class:"tag-quote tag-follow"},le={key:0,class:"options"},ie=$({__name:"tag-item",props:{tag:{},showAction:{type:Boolean},checkFollowing:{type:Boolean}},setup(T){const t=T,r=i(!1),m=q(()=>{let e=[];return t.tag.is_following===0?e.push({label:"关注",key:"follow"}):(t.tag.is_top===0?e.push({label:"置顶",key:"stick"}):e.push({label:"取消置顶",key:"unstick"}),e.push({label:"取消关注",key:"unfollow"})),e}),l=e=>{switch(e){case"follow":I({topic_id:t.tag.id}).then(o=>{t.tag.is_following=1,window.$message.success("关注成功")}).catch(o=>{console.log(o)});break;case"unfollow":z({topic_id:t.tag.id}).then(o=>{t.tag.is_following=0,window.$message.success("取消关注")}).catch(o=>{console.log(o)});break;case"stick":F({topic_id:t.tag.id}).then(o=>{t.tag.is_top=o.top_status,window.$message.success("置顶成功")}).catch(o=>{console.log(o)});break;case"unstick":F({topic_id:t.tag.id}).then(o=>{t.tag.is_top=o.top_status,window.$message.success("取消置顶")}).catch(o=>{console.log(o)});break}};return A(()=>{r.value=!1}),(e,o)=>{const w=x("router-link"),g=H,k=B,a=J,d=P,v=Q,p=R;return!e.checkFollowing||e.checkFollowing&&e.tag.is_following===1?(c(),_("div",se,[n(p,null,{header:s(()=>[(c(),b(k,{type:"success",size:"large",round:"",key:e.tag.id},{avatar:s(()=>[n(g,{src:e.tag.user.avatar},null,8,["src"])]),default:s(()=>[n(w,{class:"hash-link",to:{name:"home",query:{q:e.tag.tag,t:"tag"}}},{default:s(()=>[L(" #"+f(e.tag.tag),1)]),_:1},8,["to"]),e.showAction?u("",!0):(c(),_("span",ae,"("+f(e.tag.quote_num)+")",1)),e.showAction?(c(),_("span",ce,"("+f(e.tag.quote_num)+")",1)):u("",!0)]),_:1}))]),"header-extra":s(()=>[e.showAction?(c(),_("div",le,[n(v,{placement:"bottom-end",trigger:"click",size:"small",options:m.value,onSelect:l},{default:s(()=>[n(d,{type:"success",quaternary:"",circle:"",block:""},{icon:s(()=>[n(a,null,{default:s(()=>[n(h(U))]),_:1})]),_:1})]),_:1},8,["options"])])):u("",!0)]),_:1})])):u("",!0)}}});const _e=$({__name:"Topic",setup(T){const t=ne(),r=i([]),m=i("hot"),l=i(!1),e=i(!1),o=i(!1);D(e,()=>{e.value||(window.$message.success("保存成功"),t.commit("refreshTopicFollow"))});const w=q({get:()=>{let a="编辑";return e.value&&(a="保存"),a},set:a=>{}}),g=()=>{l.value=!0,j({type:m.value,num:50}).then(a=>{r.value=a.topics,l.value=!1}).catch(a=>{console.log(a),l.value=!1})},k=a=>{m.value=a,a=="follow"?o.value=!0:o.value=!1,g()};return A(()=>{g()}),(a,d)=>{const v=te,p=X,C=B,V=Z,M=ie,N=ee,O=oe,S=W;return c(),_("div",null,[n(v,{title:"话题"}),n(S,{class:"main-content-wrap tags-wrap",bordered:""},{default:s(()=>[n(V,{type:"line",animated:"","onUpdate:value":k},K({default:s(()=>[n(p,{name:"hot",tab:"热门"}),n(p,{name:"new",tab:"最新"}),h(t).state.userLogined?(c(),b(p,{key:0,name:"follow",tab:"关注"})):u("",!0)]),_:2},[h(t).state.userLogined?{name:"suffix",fn:s(()=>[n(C,{checked:e.value,"onUpdate:checked":d[0]||(d[0]=y=>e.value=y),checkable:""},{default:s(()=>[L(f(w.value),1)]),_:1},8,["checked"])]),key:"0"}:void 0]),1024),n(O,{show:l.value},{default:s(()=>[n(N,null,{default:s(()=>[(c(!0),_(Y,null,G(r.value,y=>(c(),b(M,{tag:y,showAction:h(t).state.userLogined&&e.value,checkFollowing:o.value},null,8,["tag","showAction","checkFollowing"]))),256))]),_:1})]),_:1},8,["show"])]),_:1})])}}});const Me=E(_e,[["__scopeId","data-v-15794a53"]]);export{Me as default};
