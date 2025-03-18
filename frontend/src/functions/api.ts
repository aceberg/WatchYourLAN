const api = 'http://0.0.0.0:8840';

export const apiGetAllHosts = async () => {
  const url = api+'/api/all';
  const hosts = await (await fetch(url)).json();

  return hosts;
};

export const apiEditHost = async (id:number, name:string, known:string) => {

  const url = api+'/api/edit/'+id+'/'+name+'/'+known;
  const res = await (await fetch(url)).json();

  return res;
};

export const apiGetHost = async (id:string) => {

  const url = api+'/api/host/'+id;
  const res = await (await fetch(url)).json();

  return res;
};

export const apiDelHost = async (id:number) => {

  const url = api+'/api/host/del/'+id;
  const res = await (await fetch(url)).json();

  return res;
};

export const apiPortScan = async (ip:string, port:number) => {

  const url = api+'/api/port/'+ip+'/'+port;
  const res = await (await fetch(url)).json();

  return res;
};