import{u as P,b as R}from"./vue-router-e5a2430e.js";import{b as S}from"./formatTime-4210fcd1.js";import{d as k,e as o,f as s,j as e,k as a,x as l,bf as b,H as c,b as U,w as f,Y as h,F as y,u as V,q as x}from"./@vue-a481fc63.js";import{o as M,F as T,Q as j,I as E,G}from"./naive-ui-d8de3dda.js";import{_ as w,b as H}from"./index-4afa107a.js";import{_ as L}from"./post-skeleton-649e932a.js";import{_ as Q}from"./main-nav.vue_vue_type_style_index_0_lang-088e2761.js";import{u as Y}from"./vuex-44de225f.js";import"./moment-2ab8298d.js";import"./seemly-76b7b838.js";import"./vueuc-39372edb.js";import"./evtd-b614532e.js";import"./@css-render-7124a1a5.js";import"./vooks-6d99783e.js";import"./vdirs-b0483831.js";import"./@juggle-41516555.js";import"./css-render-6a5c5852.js";import"./@emotion-8a8e73f6.js";import"./lodash-es-8412e618.js";import"./treemate-25c27bff.js";import"./async-validator-dee29e8b.js";import"./date-fns-975a2d8f.js";import"./axios-4a70c6fc.js";import"./@vicons-9939c40b.js";/* empty css               */const A={class:"avatar"},J={class:"base-info"},K={class:"username"},O={class:"user-info"},W={class:"info-item"},X={class:"info-item"},Z=k({__name:"contact-item",props:{contact:{}},setup(C){const u=P(),m=t=>{u.push({name:"user",query:{s:t}})};return(t,n)=>{const _=M;return o(),s("div",{class:"contact-item",onClick:n[0]||(n[0]=i=>m(t.contact.username))},[e("div",A,[a(_,{size:54,src:t.contact.avatar},null,8,["src"])]),e("div",J,[e("div",K,[e("strong",null,l(t.contact.nickname),1),e("span",null," @"+l(t.contact.username),1)]),e("div",O,[e("span",W,"UID. "+l(t.contact.user_id),1),e("span",X,l(b(S)(t.contact.created_on))+" 加入",1)])])])}}});const tt=w(Z,[["__scopeId","data-v-644d2c15"]]),et={key:0,class:"skeleton-wrap"},ot={key:1},nt={key:0,class:"empty-wrap"},st={key:0,class:"pagination-wrap"},at=k({__name:"Contacts",setup(C){const u=Y(),m=R(),t=c(!1),n=c([]),_=c(+m.query.p||1),i=c(20),d=c(0),$=r=>{_.value=r,v()};U(()=>{v()});const v=(r=!1)=>{n.value.length===0&&(t.value=!0),H({page:_.value,page_size:i.value}).then(p=>{t.value=!1,n.value=p.list,d.value=Math.ceil(p.pager.total_rows/i.value),r&&setTimeout(()=>{window.scrollTo(0,99999)},50)}).catch(p=>{t.value=!1})};return(r,p)=>{const I=Q,z=L,B=E,q=tt,D=G,F=T,N=j;return o(),s(y,null,[e("div",null,[a(I,{title:"好友"}),a(F,{class:"main-content-wrap",bordered:""},{default:f(()=>[t.value?(o(),s("div",et,[a(z,{num:i.value},null,8,["num"])])):(o(),s("div",ot,[n.value.length===0?(o(),s("div",nt,[a(B,{size:"large",description:"暂无数据"})])):h("",!0),(o(!0),s(y,null,V(n.value,g=>(o(),x(D,{key:g.user_id},{default:f(()=>[a(q,{contact:g},null,8,["contact"])]),_:2},1024))),128))]))]),_:1})]),d.value>0?(o(),s("div",st,[a(N,{page:_.value,"onUpdate:page":$,"page-slot":b(u).state.collapsedRight?5:8,"page-count":d.value},null,8,["page","page-slot","page-count"])])):h("",!0)],64)}}});const Nt=w(at,[["__scopeId","data-v-3b2bf978"]]);export{Nt as default};
