"use client";

import { send } from 'process';
import { useState } from 'react';

import { sendEmail } from '@/components/ui/verification';

export default function EmailVerificationForm() {
  
  
  fetch("http://localhost:5656/world", {
    method: "POST",
    mode: 'no-cors',
    headers: {
      "Content-Type": "application/json"
    },
    body: JSON.stringify({
      NameOfWorld: "testworld" ,
      NumberOfUsers: "testworld" ,
      Map: "testworld",
      AccountEmail: localStorage.getItem('email')
    })

  }).then(() => {
    console.log('World data sent');
   
  }
  ).catch((error) => {
    console.error('Failed to send world data', error);
  });


  return (

    <div>  
      <button>Submit</button>
    </div>
  )
}
