import { StrictMode } from "react";
import { createRoot } from "react-dom/client";
import { MantineProvider } from "@mantine/core";
import { theme } from "./config";
import { createRouter, RouterProvider } from "@tanstack/react-router";
import { routeTree } from "./routeTree.gen";

import "@mantine/core/styles.css";
import "./style.css";

const router = createRouter({ routeTree });

declare module "@tanstack/react-router" {
    interface Register {
        router: typeof router;
    }
}

const root = document.getElementById("root")!;
createRoot(root).render(
    <StrictMode>
        <MantineProvider theme={theme} forceColorScheme="dark">
            <RouterProvider router={router} />
        </MantineProvider>
    </StrictMode>,
);
