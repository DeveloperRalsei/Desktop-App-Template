import { useEffect, useState } from "react";
import { Outlet, createRootRoute } from "@tanstack/react-router";
import { AppShell, Burger, Group } from "@mantine/core";
import { main } from "#/models";
import { GetAppConstants } from "#/main/App";

export const Route = createRootRoute({
    component: RootComponent,
});

function RootComponent() {
    const [navbarCollapsed, setNavbarCollapsed] = useState(true);
    const [appInfo, setAppInfo] = useState<main.AppConfig | null>(null);

    useEffect(() => {
        GetAppConstants()
            .then((config) => setAppInfo(config))
            .catch((err) => console.error(err));
    }, []);

    return (
        <AppShell
            header={{
                height: 60,
                offset: true,
            }}
            navbar={{
                breakpoint: "sm",
                width: 300,
                collapsed: {
                    desktop: true,
                    mobile: navbarCollapsed,
                },
            }}
        >
            <AppShell.Header>
                <Group w="100%" h="100%" align="center" justify="space-between">
                    <Group>
                        <Burger
                            opened={!navbarCollapsed}
                            hiddenFrom="sm"
                            onClick={() => {
                                setNavbarCollapsed((p) => !p);
                            }}
                        />
                    </Group>
                    {appInfo?.name}
                    <div></div>
                </Group>
            </AppShell.Header>
            <AppShell.Navbar>Navbar</AppShell.Navbar>

            <AppShell.Main>
                <Outlet />
            </AppShell.Main>
        </AppShell>
    );
}
