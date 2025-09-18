# Lógica de Negocio y Modelos de Datos

Este documento describe la lógica de negocio principal del sistema de gestión del Club Norte y las relaciones entre los diferentes modelos de datos. El objetivo es proporcionar una comprensión clara de cómo interactúan las entidades para cumplir con los requerimientos funcionales.

## 1. Modelos Principales y sus Relaciones

A continuación, se desglosan las entidades centrales del sistema y la forma en que se conectan entre sí.

### a. Gestión de Acceso y Personal

-   **`User`**: Representa a un empleado o miembro del personal. Contiene información de identificación (nombre, email, usuario) y credenciales de acceso. La contraseña se hashea automáticamente antes de guardarse.
-   **`Role`**: Define los permisos y el nivel de acceso de un usuario (ej: 'admin', 'vendedor', 'repositor').
-   **`PointSale` (Punto de Venta)**: Representa una ubicación física donde se realizan transacciones (ej: 'Cantina', 'Secretaría').

**Relaciones:**
-   `User` (muchos) -> `Role` (uno): Cada usuario tiene un único rol, pero un rol puede ser asignado a muchos usuarios.
-   `User` (muchos) <-> `PointSale` (muchos): Un usuario puede estar asignado a múltiples puntos de venta, y un punto de venta puede tener múltiples usuarios asignados.

### b. Inventario y Productos

-   **`Product`**: Es un artículo que se puede vender. Tiene un código, nombre, precio y está asociado a una categoría.
-   **`Category`**: Agrupa productos de naturaleza similar (ej: 'Bebidas', 'Snacks', 'Artículos Deportivos').
-   **`StockDeposit`**: Representa el stock de un producto en el **depósito central**. Hay una única entrada por producto.
-   **`StockPointSale`**: Representa el stock de un producto en un **punto de venta específico**.
-   **`MovementStock`**: Es un registro de auditoría. Cada vez que el stock de un producto se mueve (del depósito a un punto de venta, o viceversa), se crea un registro aquí.

**Relaciones:**
-   `Product` (muchos) -> `Category` (uno): Un producto pertenece a una sola categoría.
-   `Product` (uno) -> `StockDeposit` (uno): Cada producto tiene un registro de stock en el depósito.
-   `Product` (uno) -> `StockPointSale` (muchos): Un producto puede tener stock en múltiples puntos de venta.
-   `MovementStock` se relaciona con `User` (quién hizo el movimiento) y `Product` (qué se movió). Los campos `From` y `To` indican el origen y destino (depósito o punto de venta).

### c. Transacciones y Finanzas

-   **`Register` (Caja)**: Modela una sesión de caja en un `PointSale`. Se registra quién la abre (`UserOpen`), con cuánto dinero (`OpenAmount`), y quién la cierra (`UserClose`) con su monto final (`CloseAmount`).
-   **`Income` (Ingreso)**: Registra una venta. Contiene el total, el método de pago y se asocia a un `Register`, un `User` y un `PointSale`.
-   **`IncomeItem`**: Detalla los productos vendidos dentro de una transacción de `Income`.
-   **`Expense` (Gasto)**: Registra una salida de dinero de una caja. Se asocia a un `Register`, `User` y `PointSale`.
-   **`IncomeSportsCourts` (Ingreso por Canchas)**: Un tipo de ingreso especializado para el alquiler de canchas deportivas.

**Relaciones:**
-   `Register` (muchos) -> `PointSale` (uno): Cada apertura de caja pertenece a un único punto de venta.
-   `Register` se asocia con dos usuarios: `UserOpen` y `UserClose`.
-   `Income` y `Expense` (muchos) -> `Register` (uno): Todos los ingresos y gastos se registran contra una caja abierta.
-   `Income` (uno) -> `IncomeItem` (muchos): Una venta puede incluir múltiples productos.
-   `IncomeSportsCourts` se relaciona con `SportsCourt` (la cancha alquilada) y `User` (quién registró el alquiler).

## 2. Flujos de Lógica de Negocio Clave

### Flujo 1: Realizar una Venta en un Punto de Venta

1.  Un `User` con el `Role` 'vendedor' inicia sesión y abre una `Register` (caja) en su `PointSale` asignado, registrando el `OpenAmount`.
2.  El cliente solicita uno o más `Product`s.
3.  El sistema crea un registro de `Income` asociado a la `Register` actual.
4.  Por cada producto vendido, se crea un `IncomeItem` que detalla la cantidad y el precio.
5.  El stock del producto se descuenta del `StockPointSale` correspondiente al `PointSale` de la venta.
6.  El total de la venta se suma a los acumuladores del `Register`.

### Flujo 2: Mover Stock del Depósito a un Punto de Venta

1.  Un `User` con el `Role` 'repositor' inicia el proceso.
2.  Se crea un `MovementStock`.
    -   `FromType` es 'deposit', `FromID` es el ID del depósito (o un valor que lo represente).
    -   `ToType` es 'point_sale', `ToID` es el ID del `PointSale` destino.
    -   Se especifica el `ProductID` y la `Amount` (cantidad).
3.  El sistema actualiza los registros de stock:
    -   Resta la `Amount` del `StockDeposit` del producto.
    -   Suma la `Amount` al `StockPointSale` del producto en el punto de venta destino.

### Flujo 3: Cierre de Caja

1.  Al final del turno, el `User` realiza el cierre de caja.
2.  El sistema calcula los totales de `Income` y `Expense` (separados por efectivo y otros métodos) para la `Register` abierta.
3.  El usuario cuenta el dinero y lo ingresa como `CloseAmount`.
4.  La `Register` se marca como cerrada (`IsClose = true`) y se asocia el `UserCloseID`.
5.  El sistema puede calcular y mostrar cualquier diferencia entre el dinero esperado y el `CloseAmount` real.
