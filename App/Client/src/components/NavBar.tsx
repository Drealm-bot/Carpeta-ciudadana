import { useEffect, useState } from "react";
import { Navbar, Container, Stack, Button, Modal } from "react-bootstrap" 
import { Signup } from "./SignupForm";
export const NavBar = () => {
  const [openModal, setOpenModal] = useState(false)
  const handleRegister = () => {
    return(
      <Signup/>
    )
  }

    return(
      <>
      <Navbar color="light-gray" expand="lg">
        <Container>
          <Navbar.Brand href="/">Carpeta Ciudadana</Navbar.Brand>
          <Navbar.Toggle aria-controls="basic-navbar-nav" />
          <Stack direction="horizontal" gap={2} >
            <Button as="a" variant="primary">
              Ingreso
            </Button>
            <Button onClick={() => setOpenModal(true)}>
              Registro
            </Button>
          </Stack>
        </Container>
      </Navbar>
      <Modal open={openModal} onClose={() => setOpenModal(false)}>
        Hola
      </Modal>
      </>
    )
}    
