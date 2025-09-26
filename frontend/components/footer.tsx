"use client";
import React from 'react'
import Link from 'next/link';

const Footer = () => {
  return (
    <footer>
      <div className="mx-auto max-w-5xl px-4 py-4 sm:px-6 lg:px-8">
          {/* <Link href="/">
              <img src="public/EPInterMap.png" alt=""></img>
          </Link> */}
          <div>
              <ul className="mt-4 flex flex-wrap justify-center text-white gap-6 md:gap-8 lg:gap-12 font-semibold">
                  <li>
                      <Link className="transition hover:text-[#0091d3]" href="#">Ressources</Link>
                  </li>
                  <li>
                      <Link className="transition hover:text-[#0091d3]" href="#">Explore</Link>
                  </li>
                  <li>
                      <Link className="transition hover:text-[#0091d3]" href="#">FAQ</Link>
                  </li>
              </ul>
          </div>
      </div>
      <p className="text-xs text-white text-center">&copy; 2025. EPInterMap. All rights reserved.</p>
    </footer>
  )
}

export default Footer