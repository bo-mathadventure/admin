"use client"
import React, { useState } from 'react'
import FormLoading from "@/components/ui/loading"
import FirstForm from "@/components/ui/firstform"
import formData1 from "@/components/ui/firstform"




const Page = () => {
return (
    <div>   
        <div className="signUpWrapper ">
            <div id="form1" className={`p-4  ${formData1 != null ? '' : 'hidden'}`}>
                <FirstForm  />
            </div>
            <div id="form3" className='p-4 hidden'> 
                <FormLoading/>
            </div>  
        </div>
    </div>
)
}

export default Page;