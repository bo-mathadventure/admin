import React, { useState } from 'react'
import * as z from "zod"
import { Button } from "@/components/ui/button"
import { zodResolver } from "@hookform/resolvers/zod"
import { useForm } from 'react-hook-form'
import {
    Form,
    FormControl,
    FormField,
    FormItem,
    FormLabel,
    FormMessage,
} from "@/components/ui/form"
import { Input } from "@/components/ui/input"
import Link from 'next/link'
import NextCrypto from 'next-crypto';
import { sendEmail } from './verification'

const secretKey = process.env.SECRET_KEY || ""; // Provide a default value for SECRET_KEY if it is undefined
const crypto = new NextCrypto(secretKey);


const SignUpLayout = z.object({
    firstname : z.string()
        .min(3, "Der Vorname sollte mindestens 3 Zeichen haben.")
        .max(30, "Der Vorname sollte höchstens 30 Zeich en haben."),
    lastname : z.string()
        .min(3, "Der Nachname sollte mindestens 3 Zeichen haben.")
        .max(30, "Der Nachname sollte höchstens 30 Zeichen haben."),
    email: z.string()
        .email("Ungültige E-Mail-Adresse")
        .regex(/^[a-zA-Z0-9._%+-]+@(stud\.)?hs-bochum\.de$/, "Die E-Mail-Adresse muss eine gültige hs-bochum.de-Adresse sein"),
    password : z.string()
        .min(8, "Das Passwort sollte mindestens 8 Zeichen haben."),
    confirmPassword : z.string()
        .min(8, "Das Passwort sollte mindestens 8 Zeichen haben.")
}).refine((data) => data.password === data.confirmPassword,
    { message: "Die Passwörter stimmen nicht überein", path: ["confirmPassword"]})


localStorage.clear();

const FirstForm = () => {
    const form = useForm<z.infer<typeof SignUpLayout>>({
        resolver: zodResolver(SignUpLayout),
        defaultValues: {
            firstname: "",
            lastname: "",
            email: "",
            password: "",
            confirmPassword: ""
        }
    });
   
    async function onSubmit(data: z.infer<typeof SignUpLayout>) {
        const encrypted = await crypto.encrypt(data.password);
        try {
            await fetch("http://localhost:5656/account", {
              method: "POST",
              mode: 'no-cors',
              headers: {
                "Content-Type": "application/json"
              },
              body: JSON.stringify({
                FirstName: data.firstname,
                LastName: data.lastname,
                Email: data.email,
                EncryptedPassword: encrypted
              })
            }).then(() => {
            fetch("http://localhost:5656/verification", {
              method: "POST",
              mode: 'no-cors',
              headers: {
                "Content-Type": "application/json"
              },
              body: JSON.stringify({ email: localStorage.getItem('email') })
            });
          });
          sendEmail(data.email);                     
    } catch (error) {
      console.error('Failed to submit data', error);
    }
    document.getElementById("form1")?.classList.add("hidden");
    document.getElementById('form3')?.classList.remove('hidden');
    
} 
return (

            
    <div className="formWrapper ">
                <div className="left">
                    <h3 className="title">Willkommen zurück!</h3>
                    <p>Um mit uns in Verbindung zu bleiben, melden Sie sich bitte mit Ihren persönlichen Informationen an.</p>
                    <Link href={"/signin"}>
                        <Button className='border-zinc-500 text-white-300 hover:border-zinc-200 hover:text-white-100 hover:bg-teal-600 transition-colors border rounded-full px-8 bg-teal-600'>Sign In</Button>
                    </Link>
                </div>
                <div className="right">
                    <h3 className='text-center text-2xl font-semibold'>Hier registrieren</h3>
                    <Form {...form}>
                        <form onSubmit={form.handleSubmit(onSubmit)}>
                            <FormField
                                control={form.control}
                                name="firstname"
                                render={({ field }) => (
                                    <FormItem className='space-y-0 mb-2'>
                                        <FormLabel>Vorname</FormLabel>
                                        <FormControl>
                                            <Input placeholder="Vorname" {...field} />
                                        </FormControl>
                                        <FormMessage />
                                    </FormItem>
                                )}
                            />
                             <FormField
                                control={form.control}
                                name="lastname"
                                render={({ field }) => (
                                    <FormItem className='space-y-0 mb-2'>
                                        <FormLabel>Nachname</FormLabel>
                                        <FormControl>
                                            <Input placeholder="Nachname" {...field} />
                                        </FormControl>
                                        <FormMessage />
                                    </FormItem>
                                )}
                            />
                            <FormField
                                control={form.control}
                                name="email"
                                render={({ field }) => (
                                    <FormItem className='space-y-0 mb-2'>
                                        <FormLabel>Akademische E-Mail</FormLabel>
                                        <FormControl>
                                            <Input placeholder="example@(stud.)hs-bochum.de" {...field} />
                                        </FormControl>
                                        <FormMessage />
                                    </FormItem>
                                )}
                            />
                            <FormField
                                control={form.control}
                                name="password"
                                render={({ field }) => (
                                    <FormItem className='space-y-0 mb-2'>
                                        <FormLabel>Passwort</FormLabel>
                                        <FormControl>
                                            <Input placeholder="********" type='password' {...field} />
                                        </FormControl>
                                        <FormMessage />
                                    </FormItem>
                                )}
                            />
                            <FormField
                                control={form.control}
                                name="confirmPassword"
                                render={({ field }) => (
                                    <FormItem className='space-y-0 mb-2'>
                                        <FormLabel>Passwort bestätigen</FormLabel>
                                        <FormControl>
                                            <Input placeholder="********" type='password' {...field} />
                                        </FormControl>
                                        <FormMessage />
                                    </FormItem>
                                )}
                            />
                            <Button type="submit" className='w-full bg-teal-600 text-zinc-200  hover:text-white  hover:bg-teal-600 '  >Submit</Button>
                        </form>
                    </Form>
                </div>
            </div>
)

}

export default FirstForm;