import { css } from "@linaria/core";
import { onCleanup } from "solid-js";

import bedroomUrl from "./assets/bedroom.png";
import alarmMp3   from "./assets/alarm.mp3";  
import { useNavigate } from "@solidjs/router";

const wrap = css`
  position: relative;
  width: min(100vw, 1200px);
  aspect-ratio: 4 / 3; 
  margin: 0 auto;
  background: #000 url(${bedroomUrl}) center/cover no-repeat;
  border-radius: 16px;
  overflow: hidden;
  box-shadow: 0 10px 30px rgba(0,0,0,.35);
`;

const navBar = css`
  position: absolute;
  inset: 12px 12px auto 12px;
  display: flex; gap: 12px;
  a { 
	color: #fff; 
	text-decoration: none; 
	background: rgba(0,0,0,.35); 
	padding: 6px 10px; 
	border-radius: 8px; 
  }
`;

const hotspot = css`
  position: absolute;
  top: 16%;
  right: 30%;
  width: 10%;
  height: 10%;
  cursor: pointer;
  border-radius: 30px;

  /* reset default button styles */
  /* background: transparent; */
  background: rgba(255, 255, 255, 0.25);

  border: none;
  padding: 0;
  margin: 0;

  /* optional: for accessibility when focused */
  &:focus {
    outline: 2px solid rgba(255, 255, 255, 0.5);
    outline-offset: 2px;
  }
`;

const srOnly = css`
  position: absolute; width: 1px; height: 1px; padding: 0; margin: -1px; overflow: hidden;
  clip: rect(0,0,0,0); white-space: nowrap; border: 0;
`;

export const Front = () => {
  let audio: HTMLAudioElement | undefined;
const navigate = useNavigate(); 

  const startGame = async () => {
    try {
      audio && (audio.currentTime = 0, await audio.play());
    } catch {}
    // navigate("/game");
  };

  // create the audio lazily so it doesn't block render
  const a = new Audio(alarmMp3);
  a.preload = "auto"; a.volume = 0.5;
  audio = a;
  onCleanup(() => { audio?.pause(); audio = undefined; });

  return (
    <div class={wrap} role="img" aria-label="Dragon's bedroom scene">
      <nav class={navBar}>
        <a href="/">Home</a>
        <a href="/game">Skip intro</a>
      </nav>

      {/* Alarm clock hotspot */}
      <button class={hotspot} onClick={startGame} aria-label="Wake up! Start game">
        <span class={srOnly}>Click the alarm clock to start</span>
      </button>
    </div>
  );
}
