import{c,o as l,e as f,t as m,b as u,s as v,v as w,F as y,w as b}from"./index.js";var d=m("<i>");function I(n){const[r,i]=c([]);return l(async()=>{let t=[];t=await b(n.mac),t!=null&&(t.sort((a,e)=>a.Date<e.Date?1:-1),n.show>0&&(t=t.slice(0,n.show)),i(t))}),f(y,{get each(){return r()},children:t=>(()=>{var a=d();return u(e=>{var o="Date:"+t.Date+`
Iface:`+t.Iface+`
IP:`+t.IP+`
Known:`+t.Known,s=t.Now===0?"my-box-off":"my-box-on";return o!==e.e&&v(a,"title",e.e=o),s!==e.t&&w(a,e.t=s),e},{e:void 0,t:void 0}),a})()})}export{I as M};
