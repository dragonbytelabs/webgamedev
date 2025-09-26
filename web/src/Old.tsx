import { css } from "@linaria/core";
import { createSignal } from "solid-js";

const content = css`
  background-color: #bada55;
`;

export const Old = () => {
	const [count, setCount] = createSignal(0);

	return (
		<div class={content}>
			<div>
				<a href="/">home</a>
				<a href="/game">game</a>
			</div>
			<h1>Vite + Solid</h1>
			<h1>Vite + Solid</h1>
			<h1>Vite + Solid</h1>
			<h1>Vite + Solid</h1>
			<h1>Vite + Solid</h1>
			<div class="card">
				<button onClick={() => setCount((count) => count + 1)}>
					count is {count()}
				</button>
				<p>
					Edit <code>src/App.tsx</code> and save to test HMR
				</p>
			</div>
			<p class="read-the-docs">
				Click on the Vite and Solid logos to learn more
			</p>
		</div>
	);
};
