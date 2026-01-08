import { GlobalWindow } from 'happy-dom';
import { expect } from 'bun:test';
import * as matchers from '@testing-library/jest-dom/matchers';
import '@testing-library/jest-dom'; // Import types

// Setup happy-dom
const window = new GlobalWindow();
global.window = window as any;
global.document = window.document as any;
global.navigator = window.navigator as any;
global.HTMLElement = window.HTMLElement as any;
global.HTMLInputElement = window.HTMLInputElement as any;
global.HTMLButtonElement = window.HTMLButtonElement as any;
global.Node = window.Node as any;

// Add @testing-library/jest-dom matchers to bun:test
expect.extend(matchers);
