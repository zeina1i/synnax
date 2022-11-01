import { describe, expect, it, vi } from "vitest";
import { closeWindow, initialState, setWindowState } from "@/state";
import { MockRuntime } from "./mock/runtime";
import { configureMiddleware, middleware } from "./middleware";
import { CurriedGetDefaultMiddleware } from "@reduxjs/toolkit/dist/getDefaultMiddleware";

const state = {
	drift: initialState,
};

describe("middleware", () => {
	describe("middleware", () => {
		describe("emitting actions", () => {
			it("should emit an action if it hasn't already been emited", () => {
				const store = { getState: () => state, dispatch: vi.fn() };
				const runtime = new MockRuntime(false);
				const mw = middleware(runtime)(store);
				mw((action) => action)({ type: "test" });
				expect(runtime.emissions).toEqual([
					{
						action: { type: "test" },
						emitter: "mock",
					},
				]);
			});
			it("should not emit an action if it has already been emited", () => {
				const store = { getState: () => state, dispatch: vi.fn() };
				const runtime = new MockRuntime(false);
				const mw = middleware(runtime)(store);
				mw((action) => action)({ type: "DA@test://test" });
				expect(runtime.emissions).toEqual([]);
			});
		});
		describe("'nexting' actions", () => {
			it("should next an action if it has not been emitted by 'self' ", () => {
				const store = { getState: () => state, dispatch: vi.fn() };
				const runtime = new MockRuntime(false);
				const mw = middleware(runtime)(store);
				const next = vi.fn();
				mw(next)({ type: "test" });
				expect(next).toHaveBeenCalledWith({ type: "test" });
			});
			it("should not next an action if it has been emitted by 'self' ", () => {
				const store = { getState: () => state, dispatch: vi.fn() };
				const runtime = new MockRuntime(false);
				const mw = middleware(runtime)(store);
				const next = vi.fn();
				mw(next)({ type: "DA@mock://test" });
				expect(next).not.toHaveBeenCalled();
			});
		});
		describe("key assignment", () => {
			it("should auto-assign a key to a drift action when it isn't present", () => {
				const store = { getState: () => state, dispatch: vi.fn() };
				const runtime = new MockRuntime(false);
				const mw = middleware(runtime)(store);
				const next = vi.fn();
				mw(next)(setWindowState("created"));
				expect(next).toHaveBeenCalledWith({
					type: "drift/setWindowState",
					payload: {
						key: "mock",
						state: "created",
					},
				});
			});
			it("should not auto-assign a key to a drift action if it has been emitted", () => {
				const store = { getState: () => state, dispatch: vi.fn() };
				const runtime = new MockRuntime(false);
				const mw = middleware(runtime)(store);
				const next = vi.fn();
				mw(next)({
					type: "DA@test://drift/setWindowState",
					payload: {
						key: "mock",
						state: "created",
					},
				});
				expect(next).toHaveBeenCalledWith({
					type: "drift/setWindowState",
					payload: {
						key: "mock",
						state: "created",
					},
				});
			});
		});
		describe("action execution", () => {
			it("should call executeAction if the runtime is main", () => {
				const store = { getState: () => state, dispatch: vi.fn() };
				const runtime = new MockRuntime(true);
				const mw = middleware(runtime)(store);
				const next = vi.fn();
				mw(next)(closeWindow("test"));
				expect(runtime.hasClosed.includes("test")).toBe(true);
			});
			it("should not call executeAction if the runtime is not main", () => {
				const store = { getState: () => state, dispatch: vi.fn() };
				const runtime = new MockRuntime(false);
				const mw = middleware(runtime)(store);
				const next = vi.fn();
				mw(next)(closeWindow("test"));
				expect(runtime.hasClosed.includes("test")).toBe(false);
			});
			it("should not call executeAction if the action is not a drift action", () => {
				const store = { getState: () => state, dispatch: vi.fn() };
				const runtime = new MockRuntime(true);
				const mw = middleware(runtime)(store);
				const next = vi.fn();
				mw(next)({ type: "test" });
				expect(runtime.hasClosed.includes("test")).toBe(false);
			});
		});
	});
	describe("configureMiddleware", () => {
		it("should return a function that returns a middleware when an empty array is provided", () => {
			const runtime = new MockRuntime(true);
			const mwF = configureMiddleware([], runtime);
			expect(typeof mwF).toBe("function");
			expect(mwF([] as unknown as CurriedGetDefaultMiddleware<unknown>).length).toBe(1);
		});
		it("should call a middleware curry function when provided", () => {
			const runtime = new MockRuntime(true);
			const curry = vi.fn();
			const mw = configureMiddleware(() => {
				curry();
				return [];
			}, runtime);
			mw([] as unknown as CurriedGetDefaultMiddleware<unknown>);
			expect(curry).toHaveBeenCalled();
		});
	});
});
