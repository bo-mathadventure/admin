import { send } from 'process';
import { useState } from 'react';
import { sendEmail } from '@/components/ui/verification';

const FormLoading = () => {
  return (
    <div className="bg-[#2D2F48]  rounded-lg   w-5/6 mx-auto">
        <div className="flex items-center">
          <div className="flex items-center text-white relative">
            <div className="rounded-full transition duration-500 ease-in-out h-12 w-12 py-3 border-2 border-teal-600 bg-teal-600">
                <svg viewBox="0 3 20 20" fill="currentColor" className="check w-6 h-6 text-white mx-auto">
                    <path fillRule="evenodd" d="M16.707 5.293a1 1 0 01.083 1.32l-8 10a1 1 0 01-1.5.11l-4-4a1 1 0 011.41-1.41l3.293 3.293 7.293-9.293a1 1 0 011.32-.083z" clipRule="evenodd" />
                </svg>
            </div>
        <div className="absolute top-0 -ml-10 text-center mt-16 w-32 text-xs font-medium uppercase text-teal-600">Mail</div>
      </div>

      <div className="flex-auto border-t-2 transition duration-500 ease-in-out border-teal-600"></div>

      <div className="flex items-center text-white relative">
        <div className="rounded-full transition duration-500 ease-in-out h-12 w-12 py-3 border-2 bg-teal-600 border-teal-600">
          {/* Replace the checkmark svg with a number for incomplete steps */}
          <svg viewBox="0 3 20 20" fill="currentColor" className="check w-6 h-6 text-white mx-auto ">
                    <path fillRule="evenodd" d="M16.707 5.293a1 1 0 01.083 1.32l-8 10a1 1 0 01-1.5.11l-4-4a1 1 0 011.41-1.41l3.293 3.293 7.293-9.293a1 1 0 011.32-.083z" clipRule="evenodd" />
          </svg>
        </div>
        <div className="absolute top-0 -ml-10 text-center mt-16 w-32 text-xs font-medium uppercase text-teal-600">Information</div>
      </div>
      {/* ... */}

      <div className="flex-auto border-t-2 transition duration-500 ease-in-out border-teal-600"></div>

      <div className="flex items-center text-white relative">
        <div className="rounded-full transition duration-500 ease-in-out h-12 w-12 py-3 border-2 bg-teal-600 border-teal-600">
          {/* Replace the checkmark svg with a number for incomplete steps */}
          <svg viewBox="0 3 20 20" fill="currentColor" className="check w-6 h-6 text-white mx-auto ">
                    <path fillRule="evenodd" d="M16.707 5.293a1 1 0 01.083 1.32l-8 10a1 1 0 01-1.5.11l-4-4a1 1 0 011.41-1.41l3.293 3.293 7.293-9.293a1 1 0 011.32-.083z" clipRule="evenodd" />
          </svg>
        </div>
        <div className="absolute top-0 -ml-10 text-center mt-16 w-32 text-xs font-medium uppercase text-teal-600">Maps</div>
      </div>
      {/* ... */}
        <div className="flex-auto border-t-2 transition duration-500 ease-in-out border-teal-600"></div>
        {/* Example for an inactive step */}
        <div className="flex items-center text-gray-500 relative">
        <div className="rounded-full transition duration-500 ease-in-out h-12 w-12 py-3 border-2 bg-white border-gray-600 pl-4 pt-2">
            4
        </div>
        <div className="absolute top-0 -ml-10 text-center mt-16 w-32 text-xs font-medium uppercase text-white-500">Email Verification</div>
        </div>
            
    </div>
        <div className="mb-4 mt-40 mb-40 ml-40 mr-40">
            <label htmlFor="worldName" className="text-gray-300 block mb-2 text-center">
                                  Eine Mail Nachricht wurde an Ihre E-Mail-Adresse gesendet. Bitte lassen Sie diese Seite geöffnet und überprüfen Sie Ihr Postfach auf den Bestätigungslink.

            </label>
          </div>                        
      </div>
  );
}

export default FormLoading;