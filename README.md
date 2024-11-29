## Resumen
El proyecto consiste en una plataforma descentralizada basada en blockchain que permite la trazabilidad completa de productos desde su origen hasta el consumidor final, utilizando tokens digitales para representar materias primas y productos terminados. En esta versión, se utiliza Hyperledger Fabric 2.5 como la tecnología blockchain subyacente.

## Objetivo
Crear un sistema transparente, seguro y descentralizado que permita rastrear el movimiento de materias primas y productos a través de toda la cadena de suministro, garantizando la autenticidad y procedencia de los mismos, utilizando la red Hyperledger Fabric 2.5.

## Actores del Sistema

### 1. Productor (Producer)
- Responsable del ingreso de materias primas al sistema.
- Tokeniza las materias primas originales.
- Solo puede transferir a Fábricas.
- Registra información detallada sobre el origen y características de las materias primas.

### 2. Fábrica (Factory)
- Recibe materias primas de los Productores.
- Transforma materias primas en productos terminados.
- Tokeniza los productos terminados.
- Solo puede transferir a Minoristas.
- Registra información sobre el proceso de transformación.

### 3. Minorista (Retailer)
- Recibe productos terminados de las Fábricas.
- Distribuye productos a los consumidores finales.
- Solo puede transferir a Consumidores.
- Gestiona el inventario de productos terminados.

### 4. Consumidor (Consumer)
- Punto final de la cadena de suministro.
- Recibe productos de los Minoristas.
- Puede verificar toda la trazabilidad del producto.

## Funcionalidades Clave

### 1. Gestión de Identidad
- Cada participante se identifica mediante una dirección única en la red Hyperledger Fabric.
- El administrador del sistema registra y valida los participantes.
- Control de acceso basado en roles.
- Autenticación mediante credenciales de Hyperledger Fabric.

### 2. Tokenización
- Materias Primas:
  * Tokens únicos para cada lote de materia prima.
  * Metadata asociada (origen, características, certificaciones).
  * Trazabilidad desde el origen.

- Productos:
  * Tokens únicos para productos terminados.
  * Vinculación con tokens de materias primas utilizadas.
  * Información del proceso de transformación.

### 3. Sistema de Transferencias
- Transferencias direccionales según rol.
- Sistema de aceptación/rechazo de transferencias.
- Validación automática de permisos.
- Registro inmutable de cada transferencia.
- Confirmación mediante firma digital.

### 4. Trazabilidad
- Registro completo del ciclo de vida.
- Visualización de la cadena de custodia.
- Verificación de autenticidad.
- Historia completa de transferencias.
- Registro de transformaciones.

## Arquitectura Técnica

### 1. Frontend
- Framework: Next.js
- Características:
  * Interfaz responsive.
  * Paneles específicos por rol.
  * Integración con credenciales de Hyperledger Fabric.
  * Visualización de datos en tiempo real.
  * Sistema de notificaciones.

### 2. Smart Contracts
- Framework: Chaincode de Hyperledger Fabric
- Funcionalidades:
  * Gestión de roles.
  * Tokenización.
  * Sistema de transferencias.
  * Registro de eventos.
  * Validaciones de seguridad.

### 3. Integración Blockchain
- Red: Hyperledger Fabric 2.5
- Cliente: fabric-gateway
- Características:
  * Transacciones seguras.
  * Firma digital.
  * Manejo de eventos.

## Seguridad

### 1. Smart Contracts
- Validación de roles y permisos.
- Control de acceso granular.
- Prevención de ataques comunes.
- Auditoría de código.
- Tests exhaustivos.

### 2. Frontend
- Validación de inputs.
- Manejo seguro de claves.
- Protección contra ataques XSS.
- Gestión segura de sesiones.

### 3. Transaccional
- Firmas digitales.
- Verificación de transacciones.
- Sistema de respaldo.
- Logs de auditoría.

## Despliegue

### 1. Smart Contracts
- Red: Hyperledger Fabric 2.5
- Proceso de verificación.
- Documentación de direcciones.
- Gestión de versiones.

### 2. Frontend
- Plataforma: Vercel
- Configuración de dominios.
- SSL/TLS.
- Monitoreo y logs.

## Beneficios del Sistema

1. **Transparencia**
   - Trazabilidad completa.
   - Información verificable.
   - Historia inmutable.

2. **Seguridad**
   - Datos inmutables.
   - Transacciones verificadas.

