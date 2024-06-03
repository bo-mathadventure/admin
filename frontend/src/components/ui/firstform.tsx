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
import { env, send } from 'process'
import {BACKEND_URL} from '@/loadEnv';


require('dotenv').config({ path: '../../.env' });

const secretKey = process.env.SECRET_KEY || ""; // Provide a default value for SECRET_KEY if it is undefined
const crypto = new NextCrypto(secretKey);

const SignUpLayout = z.object({
    username : z.string()
        .min(3, "Der Nachname sollte mindestens 3 Zeichen haben.")
        .max(30, "Der Nachname sollte höchstens 30 Zeichen haben.")
        .regex(/^[a-zA-Z]+$/, "Der Name darf keine Zahlen enthalten."),
    email: z.string()
        .email("Ungültige E-Mail-Adresse")
        .regex(/^[a-zA-Z0-9._%+-]+@(stud\.)?hs-bochum\.de$/, "Die E-Mail-Adresse muss eine gültige hs-bochum.de-Adresse sein"),
    password : z.string()
        .min(8, "Das Passwort sollte mindestens 8 Zeichen haben."),
    confirmPassword : z.string()
        .min(8, "Das Passwort sollte mindestens 8 Zeichen haben.")
}).refine((data) => data.password === data.confirmPassword,
    { message: "Die Passwörter stimmen nicht überein", path: ["confirmPassword"]})


const FirstForm = () => {
    const form = useForm<z.infer<typeof SignUpLayout>>({
        resolver: zodResolver(SignUpLayout),
        defaultValues: {
            username: "",
            email: "",
            password: "",
            confirmPassword: ""
        }
    });
   
    async function onSubmit(data: z.infer<typeof SignUpLayout>) {
        try {
            const response = await fetch(BACKEND_URL + "/auth/register", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify({
                    email: data.email,
                    username: data.username,
                    language: "de",
                    password: data.password,
                    confirmPassword: data.password,
                }),
            });
            if (response.ok) {
                const responseData = await response.json();
                alert(responseData);
                document.getElementById("form1")?.classList.add("hidden");
                document.getElementById('form3')?.classList.remove('hidden');

                alert("verification email skipped for now");

                const loginResponse = await fetch(BACKEND_URL + "/auth/login", {
                    method: "POST",
                    headers: {
                        "Content-Type": "application/json",
                    },
                    body: JSON.stringify({
                        email: data.email,
                        password: data.password,
                    }),
                });
                if (loginResponse.ok) {
                    const loginData = await loginResponse.json();
                    alert("login token : " + loginData.token);
                    localStorage.setItem("authToken", loginData.token);
                    sendEmail(data.email, loginData.token);
                    
                } else {
                    console.error("Failed to login", loginResponse.status);
                }
            } else {
                const errorData = await response.json();
                if (errorData.message === "ERR_USER_EXISTS") {
                    alert("User already exists");
                } else {
                    console.error("Failed to submit data", response.status);
                }
            }
        } catch (error) {
            console.error('Failed to submit data', error);
        }

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
                                name="username"
                                render={({ field }) => (
                                    <FormItem className='space-y-0 mb-2'>
                                        <FormLabel>Username</FormLabel>
                                        <FormControl>
                                            <Input placeholder="Username" {...field} />
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