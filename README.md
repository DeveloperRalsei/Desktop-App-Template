# Desktop App Template

## Structure

- root/

  - frontend/

    - [Vite](https://vite.dev/)
    - [Mantine](https://mantine.dev/)
    - [Tanstack Router](https://tanstack.com/router/)

    - `Bindings`
      - "#/\*" Bindings and types folder
      - "@/\*" Frontend source folder

  - [Wails](https://wails.io/)

## Requirements

- [Go] v1.23.3+
- [PNPM](https://pnpm.io/) (Node v15+)

> If you want to use an another package manager you must change commands in the `wails.json` file

## Installition

1. Install [Go](https://go.dev/dl/)
2. Install `wails cli` with `go install github.com/wailsapp/wails/v2/cmd/wails@latest` command
3. Run `wails doctor`
4. Run `npm install` or `pnpm install` in `frontend/` folder
5. Run `wails dev` to start developement
6. Run `wails build` to build your application

## License

MIT
