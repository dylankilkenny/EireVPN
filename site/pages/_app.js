import 'bootstrap/dist/css/bootstrap.min.css';
import 'react-datepicker/dist/react-datepicker.css';
import '../static/css/index.css';
import React, { useEffect } from 'react';
import { initGA } from '../service/Analytics';

export default function MyApp({ Component, pageProps }) {
  useEffect(() => {
    initGA();
  });
  return <Component {...pageProps} />;
}
