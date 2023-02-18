import React from 'react';
import login from '../../util/api';
import './Login.css';

import { GoogleLogin } from '@react-oauth/google';
import { useDispatch, useSelector } from 'react-redux';
import { loginAction } from './authSlice';


const Login = () => {
  const auth = useSelector(state => state.auth);
  const dispatch = useDispatch();

  if(auth.loggedIn) {
    window.location.replace('/dashboard');
  }

  const responseGoogle = (response) => {
    console.log(response);
    dispatch(loginAction());
    const { credential } = response;
    window.location.replace('/dashboard');
  };

  return (
    <div className="login-container">
      <h1>Sign up / Sign in</h1>
      {/* <button className='google-btn' onClick={() => googlelogin()}>Sign in with Google</button> */}
      <GoogleLogin
        onSuccess={responseGoogle}
        onError={() => {
          console.log('Login Failed');
        }}
      />
    </div>
  );
};

export default Login;
