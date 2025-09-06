import{C as y,D as f,E as v,o as m,G as H,e as u,S as d,t as D,b as I,s as b,H as h,l as M,F as g}from"./index.js";async function w(a,r){let e=[];return r===""?e=await y(a):e=await f(a,r),e!=null?(e.sort((s,t)=>s.Date<t.Date?1:-1),e):[]}var x=D("<i>");function F(a){const[r,e]=v([]);let s;return m(async()=>{const t=await w(a.mac,a.date);e(t),s=setInterval(async()=>{const o=await w(a.mac,a.date);e(o)},6e4)}),H(()=>{clearInterval(s)}),u(g,{each:r,children:(t,o)=>u(d,{get when(){return o()<M()},get children(){var i=x();return I(n=>{var c="Date:"+t.Date+`
Iface:`+t.Iface+`
IP:`+t.IP+`
Known:`+t.Known,l=t.Now===0?"my-box-off":"my-box-on";return c!==n.e&&b(i,"title",n.e=c),l!==n.t&&h(i,n.t=l),n},{e:void 0,t:void 0}),i}})})}export{F as M};
