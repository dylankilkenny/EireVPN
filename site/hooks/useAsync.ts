import { useState, useEffect } from 'react';
import AbortController from 'node-abort-controller';
import Error from '../interfaces/error';

type PostFunc = (body: string) => Promise<any>;
type GetFunc = () => Promise<any>;
type MakeRequestFunc = GetFunc & PostFunc;

const fetchError: Error = {
  status: 0,
  code: 'ERR',
  title: 'Something Went Wrong',
  detail: 'Something Went Wrong'
};

export default function useAsync(makeRequestFunc: MakeRequestFunc) {
  const controller = new AbortController();
  const [data, setData] = useState();
  const [error, setError] = useState();
  const [loading, setloading] = useState(true);
  async function asyncCall() {
    try {
      const res = await makeRequestFunc();
      if (res.status == 200) {
        setData(res.data);
      } else {
        controller.abort();
        setError(res);
      }
    } catch (e) {
      setError(fetchError);
      controller.abort();
    } finally {
      setloading(false);
    }
  }
  useEffect(() => {
    asyncCall();
  }, []);

  return { data, loading, error };
}
