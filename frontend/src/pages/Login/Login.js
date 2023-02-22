import React, { useState } from 'react';
import { login } from '../../util/api';
import { parseJwt } from '../../util/misc'
import { loginAction } from './authSlice';
import './Login.css';

import { GoogleLogin } from '@react-oauth/google';
import { useDispatch, useSelector } from 'react-redux';



const Login = () => {
  const auth = useSelector(state => state.auth);
  const dispatch = useDispatch();

  if(auth.loggedIn) {
    window.location.replace('/dashboard');
  }

  const responseGoogle = (response) => {
    const { credential } = response;
    const data = parseJwt(credential)
    
    login(data.email, data.name, data.sub).then((success) => {
      if(success) {
        dispatch(loginAction());
        window.location.replace('/dashboard');
      }
      else {
        window.alert("Something went wrong while logging in");
      }
    });
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
