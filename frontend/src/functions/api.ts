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