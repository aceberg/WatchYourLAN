const api = 'http://0.0.0.0:8840';

export const getAllHosts = async () => {
  const url = api+'/api/all';
  const hosts = await (await fetch(url)).json();

  return hosts;
};